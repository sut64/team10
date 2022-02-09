import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
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
import TextField from "@material-ui/core/TextField";
import { BloodTypeInterface } from "../models/IBloodtype";
import { GenderInterface } from "../models/IGender";
import { JobTitleInterface } from "../models/IJobtitle";
import { PersonnelInterface } from "../models/IPersonnel";
import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

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

function CreatePersonnel() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [bloodtype, setBloodType] = useState<BloodTypeInterface[]>([]);
  const [gender, setGender] = useState<GenderInterface []>([]);
  const [jobtitle, setJobTitle] = useState<JobTitleInterface[]>([]);
  const [personnel, setPersonnel] = useState<Partial<PersonnelInterface>>(
    {}
  );

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
    const name = event.target.name as keyof typeof personnel;
    setPersonnel({
      ...personnel,
      [name]: event.target.value,
    });
  };
  
  const handleInputChangeName = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof personnel;
    const { value } = event.target;
    setPersonnel({ ...personnel, [id]: value });
  };
  const handleInputChangePersonalID = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof personnel;
    const { value } = event.target;
    setPersonnel({ ...personnel, [id]: value });
  };
  const handleInputChangeTel = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof personnel;
    const { value } = event.target;
    setPersonnel({ ...personnel, [id]: value });
  };
  const handleInputChangeAddress = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof personnel;
    const { value } = event.target;
    setPersonnel({ ...personnel, [id]: value });
  };
  const handleInputChangeSalary = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof personnel;
    const { value } = event.target;
    setPersonnel({ ...personnel, [id]: value });
  };


  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };




  const getBloodType = async () => {
    fetch(`${apiUrl}/bloodtypes`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBloodType(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getGender = async () => {
    fetch(`${apiUrl}/genders`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setGender(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getJobTitle = async () => {
    fetch(`${apiUrl}/jobtitle`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setJobTitle(res.data);
        } else {
          console.log("else");
        }
      });
  };
  

  
  useEffect(() => {
    getBloodType();
    getGender();
    getJobTitle();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    


    let data = {
        Name:       personnel.Name,
	    Personalid:  personnel.Personalid,  
	    BirthDay:    selectedDate,
	    Tel    :     personnel.Tel,
	    Address :    personnel.Address,
	    Salary :     convertType(personnel.Salary),
        GenderID : convertType(personnel.GenderID),
        BloodTypeID: convertType(personnel.BloodTypeID),
        JobTitleID: convertType(personnel.JobTitleID),
    };

    const requestOptionsPost = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/personnels`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true)
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true)
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
              บันทึกข้อมูลบุคคลากร
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
        <Grid item xs={12}>
            <p>ชื่อ</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลชื่อ"
                value={personnel.Name}
                onChange={handleInputChangeName}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>เลขบัตรประจำตัวประชาชน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Personalid"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกเลขบัตรประจำตัวประชาชน"
                value={personnel.Personalid}
                onChange={handleInputChangePersonalID}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>เบอร์โทร</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Tel"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลเบอร์โทร"
                value={personnel.Tel}
                onChange={handleInputChangeTel}
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <p>ที่อยู่</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Address"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลที่อยู่"
                value={personnel.Address}
                onChange={handleInputChangeAddress}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>เงินเดือน</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Salary"
                variant ="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกข้อมูลเงินเดือน"
                value={personnel.Salary}
                onChange={handleInputChangeSalary}
              />
            </FormControl>
          </Grid>
        <Grid item xs={3}>
            <FormControl fullWidth variant="outlined">
              <p>เพศ</p>
              <Select
                native
                value={personnel.GenderID}
                onChange={handleChange}
                inputProps={{
                  name: "GenderID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุเพศ
                </option>
                {gender.map((item: GenderInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Genders}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={3}>
            <FormControl fullWidth variant="outlined">
              <p>กรุ๊ปเลือด</p>
              <Select
                native
                value={personnel.BloodTypeID}
                onChange={handleChange}
                inputProps={{
                  name: "BloodTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุกรุ๊ปเลือด
                </option>
                {bloodtype.map((item: BloodTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.BloodType}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>อาชีพ</p>
              <Select
                native
                value={personnel.JobTitleID}
                onChange={handleChange}
                inputProps={{
                  name: "JobTitleID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาระบุอาชีพ
                </option>
                {jobtitle.map((item: JobTitleInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Job}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันเกิด</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="BirthDay"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันเกิด"
                  minDate={new Date("1900-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />

              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          
        
         
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/personnel"
              variant="contained"
            >
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
        </Grid>
      </Paper>
    </Container>
  );
}

export default CreatePersonnel;
