import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { PatientrecordInterface } from "../models/IPatientrecord";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        container: {
            marginTop: theme.spacing(2),
        },
        table: {
            minWidth: 650,
        },
        tableSpace: {
            marginTop: 20,
        },
    })
);

function Patientrecords() {
    const classes = useStyles();
    const [patientrecords, setPatientrecords] = useState<PatientrecordInterface[]>([]);
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    };

    const getPatientrecords = async () => {
        fetch(`${apiUrl}/patientrecords`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setPatientrecords(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    useEffect(() => {
        getPatientrecords();
    }, []);

    return (
        <div>
            <Container className={classes.container} maxWidth="md">
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ข้อมูลการลงทะเบียนคนไข้นอก
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/patientrecord/create"
                            variant="contained"
                            color="primary"
                        >
                            ลงทะเบียนคนไข้นอก
                        </Button>
                    </Box>
                </Box>
                <TableContainer component={Paper} className={classes.tableSpace}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center" width="5%">
                                    ลำดับ
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    คำนำหน้า
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    ชื่อ
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    นามสกุล
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    เพศ
                                </TableCell>
                                <TableCell align="center" width="25%">
                                    เลขประจำตัวประชาชน
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    อายุ
                                </TableCell>

                                <TableCell align="center" width="20%">
                                    จังหวัด
                                </TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {patientrecords.map((item: PatientrecordInterface) => (
                                <TableRow key={item.ID}>
                                    <TableCell align="center">{item.ID}</TableCell>
                                    <TableCell align="center">{item.Prename.Prename}</TableCell>
                                    <TableCell align="center">{item.Firstname}</TableCell>
                                    <TableCell align="center">{item.Lastname}</TableCell>
                                    <TableCell align="center">{item.Gender.Genders}</TableCell>
                                    <TableCell align="center">{item.Idcardnumber}</TableCell>
                                    <TableCell align="center">{item.Age}</TableCell>

                                    <TableCell align="center">{item.Province.Province}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Container>
        </div>
    );
}

export default Patientrecords;
