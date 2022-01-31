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
import { TreatmentrecordInterface } from "../models/ITreatmentrecord";
import { format } from 'date-fns';

import moment from 'moment';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 700,

    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Treatmentrecord() {
  const classes = useStyles();
  const [treatmentrecords, setTreatmentrecords] = useState<TreatmentrecordInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getTreatmentrecord = async () => {
    fetch(`${apiUrl}/treatmentrecord`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setTreatmentrecords(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getTreatmentrecord();
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
              ข้อมูลการบันทึกการรักษา
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/CreateTreatmentRecord"
              variant="contained"
              color="primary"
            >
              บันทึกการรักษา
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ID
                </TableCell>
                <TableCell align="center" width="20%">
                  ชื่อผู้ป่วย
                </TableCell>
                <TableCell align="center" width="15%">
                  โรค
                </TableCell>
                <TableCell align="center" width="10%">
                  ยา
                </TableCell>
                <TableCell align="center" width="30%">
                  วิธีการรักษา
                </TableCell>
                <TableCell align="center" width="5%">
                  อุณหภูมิ
                </TableCell>
                <TableCell align="center" width="20%">
                  ผู้ตรวจ
                </TableCell>
                <TableCell align="center" width="20%">
                  วันที่ตรวจ
                </TableCell>

              </TableRow>
            </TableHead>
            <TableBody>
              {treatmentrecords.map((item: TreatmentrecordInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Patientrecord.Firstname} {item.Patientrecord.Lastname}</TableCell>
                  <TableCell align="center">{item.Disease.Diname}</TableCell>
                  <TableCell align="center">{item.Medicine.Medname}</TableCell>
                  <TableCell align="center">{item.Treatment}</TableCell>
                  <TableCell align="center">{item.Temperature}</TableCell>
                  <TableCell align="center">{item.Personnel.Name}</TableCell>
                  <TableCell align="center">{moment(item.RecordDate).format("DDMMYYYY ")}</TableCell>




                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Treatmentrecord;
