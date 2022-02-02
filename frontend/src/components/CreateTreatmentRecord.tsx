import React from "react";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
    makeStyles,
    Theme,
    createStyles,
    alpha,
} from "@material-ui/core/styles";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { DiseaseInterface } from "../models/IDisease";
import { MedicineInterface } from "../models/IMedicine";
import { TreatmentrecordInterface } from "../models/ITreatmentrecord";
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
import { PersonnelInterface } from "../models/IPersonnel";
import {
    MuiPickersUtilsProvider,
    KeyboardDateTimePicker,
    KeyboardDatePicker
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import image from "../image/pexels-pixabay-40568.jpg";
import SendIcon from '@material-ui/icons/Save';
import './style/CreateTreatmentRecord.css';


const Alert = (props: AlertProps) => {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
            fontWeight: 700,
            fontFamily: 'monospace',
        },
        label: {
            marginBottom: '6px',
        },
        rowleap: {
            marginTop: '-25px'
        },

        container: {
            marginTop: theme.spacing(4),

        },

        information: {
            width: '65%',
            height: '60vh',
            padding: theme.spacing(5),
            backgroundColor: '#fff',
            borderRadius: '20px'
        },

        button: {
            marginTop: theme.spacing(9),
        },

        celsius: {
            position: 'absolute',
            left: 70,
            bottom: -13,
            fontSize: 20,
            fontWeight: 900,
            color: theme.palette.common.black,
        },

        hide: {
            margin: theme.spacing(3),
        }


    })
);



function CreateTreatmentRecord() {
    const classes = useStyles();
    const [disease, setDisease] = useState<DiseaseInterface[]>([]);
    const [medicine, setMedicine] = useState<MedicineInterface[]>([]);
    const [personnel, setPersonnel] = useState<PersonnelInterface[]>([]);
    const [patient, setPatient] = useState<PatientrecordInterface[]>([]);
    const [treatmentrecord, setTreatmentrecord] = useState<Partial<TreatmentrecordInterface>>(
        {}
    );
    const [selectedDate, setSelectedDate] = React.useState<Date | null>(new Date());
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: {

            Authorization: `Bearer ${localStorage.getItem("token")}`, "Content-Type": "application/json"
        },
    };

    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    }

    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>
    ) => {
        const name = event.target.name as keyof typeof treatmentrecord;
        setTreatmentrecord({
            ...treatmentrecord,
            [name]: event.target.value,
        });
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof treatmentrecord;
        const { value } = event.target;
        setTreatmentrecord({ ...treatmentrecord, [id]: value });
    };



    const handleDateChange = (date: Date | null) => {
        console.log(date);
        setSelectedDate(date);
    };

    const getPatient = async () => {
        fetch(`${apiUrl}/patientrecords`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setPatient(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    const getPersonnel = async () => {
        fetch(`${apiUrl}/personnels`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setPersonnel(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    const getDisease = async () => {
        fetch(`${apiUrl}/disease`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setDisease(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    const getMedicine = async () => {
        fetch(`${apiUrl}/medicine`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setMedicine(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    useEffect(() => {
        getPatient();
        getPersonnel();
        getDisease();
        getMedicine();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    function submit() {

        let data = {
            Treatment: treatmentrecord.Treatment,
            Temperature: convertType(treatmentrecord.Temperature),
            PatientrecordID: convertType(treatmentrecord.PatientrecordID),
            PersonnelID: convertType(treatmentrecord.PersonnelID),
            DiseaseID: convertType(treatmentrecord.DiseaseID),
            MedicineID: convertType(treatmentrecord.MedicineID),
            RecordDate: selectedDate,
        };

        const requestOptionsPost = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        };

        fetch(`${apiUrl}/treatmentrecord`, requestOptionsPost)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("");
                } else {

                    setError(true);
                    setErrorMessage(res.error);
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
                    บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
                </Alert>
            </Snackbar>

            <div className="paper">
                <Box display="flex">
                    <div className="box-image">
                        <img className="img" src={image} />
                        <div className="content_p">
                            <div>
                                <h2 className="h_content">ระบบจัดเก็บข้อมูลการรักษา</h2>
                                <p className="text_content">
                                  คือ ระบบที่จะบันทึกการรักษาและอาการของผู้ป่วย โดยจะให้ผู้ป่วยพบ
                                    แพทย์เพื่อตรวจรักษา และแพทย์จะรวบรวมข้อมูลของผู้ป่วยนั้น บันทึกในระบบ โดยจะเก็บข้อมูล ข้อมูลผู้ป่วย
                                    โรค วิธีการรักษา สั่งยา ผู้ตรวจ เพื่อจัดส่งข้อมูลนี้ให้กับแผนกที่เกี่ยวข้องต่อไป</p>
                            </div>
                        </div>
                    </div>
                    <Box className={classes.information}>
                        <Box display="flex">
                            <Box flexGrow={1}>
                                <Typography
                                    component="h3"
                                    variant="h5"
                                    color="inherit"
                                    gutterBottom
                                    className={classes.hide}
                                >

                                </Typography>
                                <Divider />
                            </Box>
                        </Box>
                        <Grid container spacing={3} className={classes.root}>
                            <Grid item xs={6} >
                                <FormControl fullWidth variant="outlined" size="small"  >
                                    <p className={classes.label}>Patient</p>
                                    <Select
                                        native
                                        className={classes.label}
                                        value={treatmentrecord.PatientrecordID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: "PatientrecordID",
                                        }}
                                    >
                                        <option aria-label="None" value="">
                                            กรุณาเลือกผู้ป่วย
                                        </option>
                                        {patient.map((item: PatientrecordInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.Firstname} {item.Lastname}
                                            </option>
                                        ))}
                                    </Select>
                                </FormControl>


                            </Grid>


                            <Grid item xs={6}>
                                <FormControl fullWidth variant="outlined">
                                    <p>Date</p>
                                    <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                        < KeyboardDateTimePicker
                                            value={selectedDate}
                                            id="RecordDate"
                                            onChange={handleDateChange}
                                            format="yyyy-MM-dd hh:mm a"
                                            KeyboardButtonProps={{
                                                "aria-label": "change date",
                                            }}
                                        />
                                    </MuiPickersUtilsProvider>
                                </FormControl>
                            </Grid>


                            <Grid item xs={6} className={classes.rowleap}>
                                <FormControl fullWidth variant="outlined" size="small"  >
                                    <p className={classes.label}>Disease</p>
                                    <Select
                                        native
                                        value={treatmentrecord.DiseaseID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: "DiseaseID",
                                        }}
                                    >
                                        <option aria-label="None" value="">
                                            กรุณาเลือกโรค
                                        </option>
                                        {disease.map((item: DiseaseInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.Diname}
                                            </option>
                                        ))}
                                    </Select>
                                </FormControl>


                            </Grid>


                            <Grid item xs={6} className={classes.rowleap}>
                                <FormControl fullWidth variant="outlined" size="small"  >
                                    <p className={classes.label}>Medicine</p>
                                    <Select
                                        native
                                        value={treatmentrecord.MedicineID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: "MedicineID",
                                        }}
                                    >
                                        <option aria-label="None" value="">
                                            กรุณาเลือกยา
                                        </option>
                                        {medicine.map((item: MedicineInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.Medname}
                                            </option>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>

                            <Grid item xs={12} className={classes.rowleap}>
                                <p className={classes.label}>Treatment</p>
                                <FormControl fullWidth variant="outlined">
                                    <TextField
                                        id="Treatment"
                                        type="string"
                                        placeholder="กรุณากรอกวิธีการรักษา"
                                        value={treatmentrecord.Treatment}
                                        onChange={handleInputChange}

                                    />
                                </FormControl>
                            </Grid>

                            <Grid item xs={6} className={classes.rowleap}>
                                <FormControl fullWidth variant="outlined" size="small"  >
                                    <p className={classes.label}>Doctor</p>
                                    <Select
                                        native
                                        value={treatmentrecord.PersonnelID}
                                        onChange={handleChange}
                                        inputProps={{
                                            name: "PersonnelID",
                                        }}
                                    >
                                        <option aria-label="None" value="">
                                            กรุณาเลือกผู้ตรวจ
                                        </option>
                                        {personnel.map((item: PersonnelInterface) => (
                                            <option value={item.ID} key={item.ID}>
                                                {item.Name}
                                            </option>
                                        ))}
                                    </Select>
                                </FormControl>
                            </Grid>


                            <Grid item xs={2} className={classes.rowleap}>
                                <p className={classes.label}>Temperature</p>
                                <FormControl fullWidth variant="outlined" >
                                    <p className={classes.celsius}>&#8451;</p>
                                    <TextField
                                        id="Temperature"
                                        type="number"
                                        variant="outlined"
                                        size="small"
                                        value={treatmentrecord.Temperature}
                                        onChange={handleInputChange}
                                    />
                                </FormControl>
                            </Grid>
                        </Grid>

                        <Grid item xs={9} className={classes.button}>
                            <Button
                                component={RouterLink}
                                to="/TreatmentRecord"
                                variant="contained"
                            >
                                ข้อมูลการบันทึกการรักษา
                            </Button>

                            <Button
                                className="submit"
                                startIcon={<SendIcon />}
                                onClick={submit}
                                variant="contained"
                                color="primary"
                            >
                                บันทึกการรักษา
                            </Button>
                        </Grid>
                    </Box>
                </Box>
            </div>
        </Container>

    );
}
export default CreateTreatmentRecord;