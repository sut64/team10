import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {makeStyles,Theme,createStyles,alpha,} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { HistorySheetInterface } from "../models/IHistorySheet";
import { DrugAllergyInterface } from "../models/IDrugAllergy";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { PersonnelInterface } from "../models/IPersonnel";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { TextField } from "@material-ui/core";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function HistorySheetCreate() {
 const classes = useStyles();
 const [selectedDate,       setSelectedDate]    = useState<Date | null>(new Date());
 const [personnels,         setPresonnel]       = useState<PersonnelInterface[]>([]);
 const [patientrecords,    setPatientrecord]  = useState<PatientrecordInterface[]>([]);
 const [drug_allergys,      setDrugAllergy]     = useState<DrugAllergyInterface[]>([]);
 const [history_sheets,     setHistorySheet]    = useState<Partial<HistorySheetInterface>>({});
 
 const [success, setSuccess] = useState(false);
 const [error, setError] = useState(false);

 const apiUrl = "http://localhost:8080";
 const requestOptions = {
   method: "GET",
   headers: {
     Authorization: `Bearer ${localStorage.getItem("token")}`,
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
 const handleInputChange = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof history_sheets;
    const { value } = event.target;
    setHistorySheet({ ...history_sheets, [id]: value });
  };

 const handleChange = (
   event: React.ChangeEvent<{ name?: string; value: unknown }>
 ) => {
   const name = event.target.name as keyof typeof history_sheets;
   setHistorySheet({
     ...history_sheets,
     [name]: event.target.value,
   });
 };

 const handleDateChange = (date: Date | null) => {
   console.log(date);
   setSelectedDate(date);
 };

 const getPersonnel = async () => {
  fetch(`${apiUrl}/category`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {setPresonnel(res.data);} 
      else {console.log("else");}
    });
  };  
  const getPatientrecord = async () => {
    fetch(`${apiUrl}/category`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {setPatientrecord(res.data);} 
        else {console.log("else");}
      });
    }; 
 const getDrugAllergy = async () => {
  fetch(`${apiUrl}/category`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {setDrugAllergy(res.data);} 
      else {console.log("else");}
    });
  }; 
   

 useEffect(() => {
     getPersonnel();
     getPatientrecord();
     getDrugAllergy();
  }, []);
  
 const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
  };  

function submit() {
   let data = {
     PersonnelID : convertType(history_sheets.PersonnelID),
     PatientrecordID : convertType(history_sheets.PatientrecordID),
     DrugAllergyID : convertType(history_sheets.DrugAllergyID),
     Weight : history_sheets.Weight ?? "",
     Height : history_sheets.Height ?? "",
     PressureOn : history_sheets.PressureOn ?? "",
     PressureLow : history_sheets.PressureLow ?? "",
     Symptom : history_sheets.Symptom ?? "",
   };

   console.log(data)

   const requestOptionsPost = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };

  fetch(`${apiUrl}/historysheet`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกผิดพลาด")
          setError(true);
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
        บันทึกข้อมูลผิดพลาด
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
             Create History Sheet
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={3} className={classes.root}>
         
       <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Email</p>
              <Select
                native
                value={history_sheets.PersonnelID}
                onChange={handleChange}
                inputProps={{
                  name: "PersonnelID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกชื่อคนไข้
                </option>
                {patientrecords.map((item: PatientrecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Prename}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ยาที่แพ้</p>
              <Select
                native
                value={history_sheets.ID}
                onChange={handleChange}
                inputProps={{
                  name: "ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกยาที่แพ้
                </option>
                {drug_allergys.map((item: DrugAllergyInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.DName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>น้ำหนัก</p>
              <TextField            
                id="weight"
                variant="outlined"
                type="float"
                size="medium"
                value={history_sheets.Weight || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>ส่วนสูง</p>
              <TextField            
                id="height"
                variant="outlined"
                type="float"
                size="medium"
                value={history_sheets.Height || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>ความดันบน</p>
              <TextField            
                id="pressureon"
                variant="outlined"
                type="uint"
                size="medium"
                value={history_sheets.PressureOn || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>ความดันล่าง</p>
              <TextField            
                id="pressurelow"
                variant="outlined"
                type="uint"
                size="medium"
                value={history_sheets.PressureLow || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>อาการเบื้องต้น</p>
              <TextField            
                id="symptom"
                variant="outlined"
                type="string"
                size="medium"
                value={history_sheets.Symptom || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

         <Grid item xs={12}>
           <Button component={RouterLink} to="/historysheets" variant="contained">
             History Sheet Table
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
           >
             Submit
           </Button>
         </Grid>
       </Grid>
     </Paper>
   </Container>
 );
}
export default HistorySheetCreate;