import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
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

import { BillInterface } from "../models/IBill";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { MedicineInterface } from "../models/IMedicine";
import { MedicalTreatmentInterface } from "../models/IMedicalTreatment";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { InputRounded, SettingsInputHdmiRounded } from "@material-ui/icons";
import { Input } from "@material-ui/core";
import { convertTypeAcquisitionFromJson } from "typescript";

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

function CreateBill() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [records, setrecords] = useState<PatientrecordInterface[]>([]);
  const [treatments, settreatments] = useState<MedicalTreatmentInterface[]>([]);
  const [medicines, setmedicines] = useState<MedicineInterface[]>([]);
  const [bills, setbills] = useState<Partial<BillInterface>>({});

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

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof bills;
    setbills({
      ...bills,
      [name]: event.target.value,
    });
  };
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof bills;
    const { value } = event.target;
    setbills({ ...bills, [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getRecord = async () => {
    fetch(`${apiUrl}/patientrecords`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setrecords(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getMedcine = async () => {
    fetch(`${apiUrl}/medicine`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setmedicines(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getTreatment = async () => {
    fetch(`${apiUrl}/medicaltreatment`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          settreatments(res.data);
        } else {
          console.log("else");
        }
      });
  };

  let lb: number = 0;

  useEffect(() => {
    getRecord();
    getMedcine();
    getTreatment();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      PatientrecordID: convertType(bills.PatientrecordID),
      Cot: convertType(bills.Cot),
      Com: convertType(bills.Com),
      listofbill: convertType(bills.Listofbill),
      Sum:  convertType(bills.Sum), 
      Dateofbill: selectedDate,
    };

    console.log(data);

    const requestOptionsPost = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/bill`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้");
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้");
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
          บันทึกข้อมูลไม่สำเร็จ
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
              บันทึกข้อมูล
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่อผู้ป่วย</p>
              <Select
                native
                value={bills.PatientrecordID}
                onChange={handleChange}
                inputProps={{
                  name: "PatientrecordID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือก
                </option>
                {records.map((item: PatientrecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Firstname} {item.Lastname}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>การรักษา</p>
              <Select
                native
                value={bills.Cot}
                onChange={handleChange}
                inputProps={{
                  name: "Cot",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกการรักษา
                </option>
                {treatments.map((item: MedicalTreatmentInterface) => (
                  <option value={item.Price} key={item.Price}>
                    {item.Tname} {item.Price} บาท
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ยา</p>
              <Select
                native
                value={bills.Com}
                onChange={handleChange}
                inputProps={{
                  name: "Com",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกการรักษา
                </option>
                {medicines.map((item: MedicineInterface) => (
                  <option value={item.Price} key={item.Price}>
                    {item.Medname} {item.Price} บาท
                  </option>
                ))}
              </Select> 
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>จำนวนรายการ</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Listofbill"
                variant ="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกข้อมูลชื่อ"
                value={bills.Listofbill}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="Dateofbill"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Button component={RouterLink} to="/Bill" variant="contained">
            กลับ
          </Button>
          <Button
            style={{ float: "right" }}
            variant="contained"
            onClick={submit}
            color="primary"
          >
            บันทึก
          </Button>
        </Grid>
        <Grid />
      </Paper>
    </Container>
  );
}

export default CreateBill;
