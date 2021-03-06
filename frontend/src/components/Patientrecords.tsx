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
            <Container className={classes.container} maxWidth="xl">
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ??????????????????????????????????????????????????????????????????????????????
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/patientrecord/create"
                            variant="contained"
                            color="primary"
                        >
                            ???????????????????????????????????????????????????
                        </Button>
                    </Box>
                </Box>
                <TableContainer component={Paper} className={classes.tableSpace}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="left" width="2%">
                                    ???????????????
                                </TableCell>
                                <TableCell align="left" width="5%">
                                    ????????????????????????
                                </TableCell>
                                <TableCell align="left" width="6%">
                                    ????????????
                                </TableCell>
                                <TableCell align="left" width="6%">
                                    ?????????????????????
                                </TableCell>
                                <TableCell align="left" width="3%">
                                    ?????????
                                </TableCell>
                                <TableCell align="left" width="11%">
                                    ??????????????????????????????????????????????????????
                                </TableCell>
                                <TableCell align="left" width="2%">
                                    ????????????
                                </TableCell>
                                <TableCell align="left" width="11%">
                                    ?????????????????????
                                </TableCell>
                                <TableCell align="center" width="6%">
                                    ??????????????????????????????
                                </TableCell>
                                <TableCell align="left" width="4%">
                                    Tel.
                                </TableCell>
                                <TableCell align="left" width="5%">
                                    Email
                                </TableCell>
                                <TableCell align="left" width="14%">
                                    ?????????????????????
                                </TableCell>
                                <TableCell align="left" width="5%">
                                    ?????????????????????
                                </TableCell>
                                <TableCell align="left" width="10%">
                                    ????????????????????????????????????????????????
                                </TableCell>
                                <TableCell align="center" width="10%">
                                    ???????????????????????????????????????????????????????????????
                                </TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {patientrecords.map((item: PatientrecordInterface) => (
                                <TableRow key={item.ID}>
                                    <TableCell align="left">{item.ID}</TableCell>
                                    <TableCell align="left">{item.Prename.Prename}</TableCell>
                                    <TableCell align="left">{item.Firstname}</TableCell>
                                    <TableCell align="left">{item.Lastname}</TableCell>
                                    <TableCell align="left">{item.Gender.Genders}</TableCell>
                                    <TableCell align="left">{item.Idcardnumber}</TableCell>
                                    <TableCell align="left">{item.Age}</TableCell>
                                    <TableCell align="left">{format((new Date(item.Birthday)), 'dd MMMM yyyy')}</TableCell>
                                    <TableCell align="center">{item.BloodType.BloodType}</TableCell>
                                    <TableCell align="left">{item.Phonenumber}</TableCell>
                                    <TableCell align="left">{item.Email}</TableCell>
                                    <TableCell align="left">{item.Home}</TableCell>
                                    <TableCell align="left">{item.Province.Province}</TableCell>
                                    <TableCell align="left">{format((new Date(item.Timestamp)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                                    <TableCell align="left">{item.Personnel.Name}</TableCell>
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
