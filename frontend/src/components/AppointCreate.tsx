import React from "react";
import { Link as RouterLink } from "react-router-dom";
import { useEffect, useState } from "react";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Select from "@material-ui/core/Select";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

import { PersonnelInterface } from "../models/IPersonnel";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { TreatmentrecordInterface } from "../models/ITreatmentrecord";
import { AppointsInterface } from "../models/IAppoint";
 
function Alert(props: AlertProps) {
 return <MuiAlert elevation={6} variant="filled" {...props} />;
}
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {flexGrow: 1},
   container: {marginTop: theme.spacing(2)},
   paper: {padding: theme.spacing(2),color: theme.palette.text.secondary},
 })
);
 
function AppointCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [personnels, setPersonnels] = useState<PersonnelInterface[]>([]);
  const [patients, setPatients] = useState<PatientrecordInterface[]>([]);
  const [treatments, setTreatments] = useState<TreatmentrecordInterface[]>([]);
  const [appoints, setAppoints] = useState<Partial<AppointsInterface>>({});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      //Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };
 
 const handleDateChange = (date: Date | null) => {
   setSelectedDate(date);
 };

const handleChange = (
  event: React.ChangeEvent<{ name?: string; value: unknown }>
) => {
  const name = event.target.name as keyof typeof appoints;
  setAppoints({
    ...appoints,
    [name]: event.target.value,
  });
};
 
 const handleInputChange = (
  event: React.ChangeEvent<{ id?: string; value: any }>
) => {
  const id = event.target.id as keyof typeof AppointCreate;
  const { value } = event.target
  setAppoints({ ...appoints, [id]: value });
};

 const getPersonnels = async () => {
  fetch(`${apiUrl}/personnels_appoint`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setPersonnels(res.data);
      } else {
        console.log("else");
      }
    });
};

const getPatients = async () => {
  fetch(`${apiUrl}/patientrecords`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setPatients(res.data);
      } else {
        console.log("else");
      }
    });
};

const getTreatments = async () => {
  fetch(`${apiUrl}/treatmentrecord`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setTreatments(res.data);
      } else {
        console.log("else");
      }
    });
};

useEffect(() => {
  getPersonnels();
  getPatients();
  getTreatments();
}, []);

const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
};

 function submit() {
   let data = {
      Appoint_ID: appoints.Appoint_ID ?? "",
      PersonnelID: convertType(appoints.PersonnelID),
      PatientrecordID: convertType(appoints.PatientrecordID),
      TreatmentrecordID: convertType(appoints.TreatmentrecordID),
      Room_number: typeof appoints.Room_number === "string" ? parseInt(appoints.Room_number) : 0,
      Date_appoint : selectedDate,
   };
   console.log(data)

   const requestOptionsPost = {
    method: "POST",
    headers: {
      //Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };
 
  fetch(`${apiUrl}/appointments`, requestOptionsPost)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      console.log("บันทึกได้")
      setSuccess(true);
      setErrorMessage("")
    } 
    else {
      console.log("บันทึกไม่ได้")
      setError(true);
      setErrorMessage(res.error)
    }
  });

 }
 
 return (
   <Container className={classes.container} maxWidth="md">
     <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="success">
         บันทึกข้อมูลสำเร็จ
       </Alert>
     </Snackbar>
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
         บันทึกข้อมูลไม่สำเร็จ {errorMessage}
       </Alert>
     </Snackbar>
     <Paper className={classes.paper}>
       <Box display="flex">
         <Box flexGrow={1}>
           <Typography
             component="h2"
             variant="h6"
             color="primary"
             gutterBottom
           >
             Create Appointment
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={3} className={classes.root}>
         <Grid item xs={6}>
           <p>Appointment ID</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="Appoint_ID"
               variant="outlined"
               type="string"
               size="medium"
               placeholder="กรอกรหัสการนัดคนไข้"
               value={appoints.Appoint_ID || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>Patient</p>
             <Select
                native
                value={appoints.PatientrecordID}
                onChange={handleChange}
                inputProps={{
                  name: "PatientrecordID",
                }}
              >
                <option aria-label="None" value="">
                  คนไข้
                </option>
                {patients.map((item: PatientrecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Firstname} {item.Lastname}
                  </option>
                ))}
              </Select>
           </FormControl>
         </Grid>
         
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>Treatment</p>
             <Select
                native
                value={appoints.TreatmentrecordID}
                onChange={handleChange}
                inputProps={{
                  name: "TreatmentrecordID",
                }}
              >
                <option aria-label="None" value="">
                  การรักษา
                </option>
                {treatments.map((item: TreatmentrecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Treatment}
                  </option>
                ))}
              </Select>
           </FormControl>
         </Grid>

         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>Doctor</p>
             <Select
                native
                value={appoints.PersonnelID}
                onChange={handleChange}
                inputProps={{
                  name: "PersonnelID",
                }}
              >
                <option aria-label="None" value="">
                  แพทย์ผู้ตรวจ
                </option>
                {personnels.map((item: PersonnelInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
           </FormControl>
         </Grid>
         <Grid item xs={4}>
           <FormControl fullWidth variant="outlined">
             <p>Room number</p>
             <TextField
               id="Room_number"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 0 ,max: 10 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={appoints.Room_number || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={8}>
           <FormControl fullWidth variant="outlined">
             <p>Date Appointment</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 id="Date_appoint"
                 format="yyyy-MM-dd"
                 value={selectedDate}
                 //minDate ={new Date()}
                 onChange={handleDateChange}
                 KeyboardButtonProps={{
                   "aria-label": "change date",
                 }}
               />
             </MuiPickersUtilsProvider>
           </FormControl>
         </Grid>
         <Grid item xs={12}>
           <Button component={RouterLink} to="/AppointTable" variant="contained">
             Back
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
           >
             บันทึกการนัด
           </Button>
         </Grid>
       </Grid>
     </Paper>
   </Container>
 );
}
 
export default AppointCreate;
