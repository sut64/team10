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
import { TextField } from "@material-ui/core";

import { HistorySheetInterface } from "../models/IHistorySheet";
import { DrugAllergyInterface } from "../models/IDrugAllergy";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { PersonnelInterface } from "../models/IPersonnel";

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
  const [personnel,         setPresonnel]    = useState<PersonnelInterface[]>([]);
  const [patientrecord,    setPatientrecord] = useState<PatientrecordInterface[]>([]);
  const [drugallergy,      setDrugAllergy]   = useState<DrugAllergyInterface[]>([]);
  const [historysheet,     setHistorySheet]  = useState<Partial<HistorySheetInterface>>({});

 const [success, setSuccess] = useState(false);
 const [error, setError] = useState(false);
 const [errorMessage, setErrorMessage] = useState("");

 const apiUrl = "http://localhost:8080";
 const requestOptions = {
   method: "GET",
   headers: { "Content-Type": "application/json" },
 };

 const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };


 const handleChange = (
   event: React.ChangeEvent<{ name?: string; value: unknown }>
 ) => {
   const name = event.target.name as keyof typeof historysheet;
   setHistorySheet({
     ...historysheet,
     [name]: event.target.value,
   });
 };

 const handleInputChangeWeight = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };
 const handleInputChangeHeight = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };
 const handleInputChangeTemperature = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };
 const handleInputChangePressureOn = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };
 const handleInputChangePressureLow = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };
 const handleInputChangeSymptom = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof historysheet;
    const { value } = event.target;
    setHistorySheet({ ...historysheet, [id]: value });
  };

  const getPersonnel = async () => {
    fetch(`${apiUrl}/personnels`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setPresonnel(res.data);
        } else {
          console.log("else");
        }
      });
    };  
    const getPatientrecord = async () => {
      fetch(`${apiUrl}/patientrecords`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
          if (res.data) {
            setPatientrecord(res.data);
          } else {
            console.log("else");
          }
        });
      }; 
   const getDrugAllergy = async () => {
    fetch(`${apiUrl}/drugallergys`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDrugAllergy(res.data);
        } else {
          console.log("else");
        }
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
    PersonnelID :      convertType(historysheet.PersonnelID),
    PatientrecordID :  convertType(historysheet.PatientrecordID),
    DrugAllergyID :    convertType(historysheet.DrugAllergyID),
    Weight :           convertType(historysheet.Weight),
    Height :           convertType(historysheet.Height),
    Temperature :      convertType(historysheet.Temperature),
    PressureOn :       convertType(historysheet.PressureOn),
    PressureLow :      convertType(historysheet.PressureLow),
    Symptom :          historysheet.Symptom,
  };

   const requestOptionsPost = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };

  fetch(`${apiUrl}/historysheets`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("")
        } else {
          console.log("บันทึกผิดพลาด")
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
        บันทึกข้อมูลผิดพลาด : {errorMessage}
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
              <p>พยาบาล</p>
              <Select
                native
                value={historysheet.PersonnelID}
                onChange={handleChange}
                inputProps={{
                  name: "PersonnelID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุพยาบาลผู้ซักประวัติ
                </option>
                {personnel.map((item: PersonnelInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>คนไข้</p>
              <Select
                native
                value={historysheet.PatientrecordID}
                onChange={handleChange}
                inputProps={{
                  name: "PatientrecordID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุชื่อคนไข้
                </option>
                {patientrecord.map((item: PatientrecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Firstname} {item.Lastname}
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
                value={historysheet.DrugAllergyID}
                onChange={handleChange}
                inputProps={{
                  name: "DrugAllergyID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุยาที่แพ้
                </option>
                {drugallergy.map((item: DrugAllergyInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name} : {item.Symptom}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <p>น้ำหนัก</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Weight"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลน้ำหนัก"
                value={historysheet.Weight}
                onChange={handleInputChangeWeight}
              />
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <p>ส่วนสูง</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Height"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลส่วนสูง"
                value={historysheet.Height}
                onChange={handleInputChangeHeight}
              />
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <p>อุณหภูมิ</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Temperature"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลอุณหภูมิ"
                value={historysheet.Temperature}
                onChange={handleInputChangeTemperature}
              />
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <p>ความดันบน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="PressureOn"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลความดันบน"
                value={historysheet.PressureOn}
                onChange={handleInputChangePressureOn}
              />
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <p>ความดันล่าง</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="PressureLow"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลความดันล่าง"
                value={historysheet.PressureLow}
                onChange={handleInputChangePressureLow}
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <p>อาการเบื้องต้น</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Symptom"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลอาการเบื้องต้น"
                value={historysheet.Symptom}
                onChange={handleInputChangeSymptom}
              />
            </FormControl>
          </Grid>

         <Grid item xs={12}>
           <Button component={RouterLink} 
           to="/historysheets" 
           variant="contained"
           >
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