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
import { BillInterface } from "../models/IBill";
import {PatientrecordInterface} from "../models/IPatientrecord";
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

function Billlist() {
  const classes = useStyles();
  const [bills, setBills] = useState<BillInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getBills = async () => {
    fetch(`${apiUrl}/bill`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setBills(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getBills();
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
              รายการใบเสร็จ
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/CreateBill"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="15%">
                  เลขบิล
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="15%">
                  นามสกุล
                </TableCell>
                <TableCell align="center" width="15%">
                  จำนวนรายการ
                </TableCell>
                <TableCell align="center" width="15%">
                  ค่ายา
                </TableCell>
                <TableCell align="center" width="15%">
                  ค่ารักษา
                </TableCell>
                <TableCell align="center" width="15%">
                  ค่าใช้จ่ายทั้งหมด
                </TableCell>
                <TableCell align="center" width="15%">
                  วันออกบิล
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bills.map((item: BillInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Patientrecord.Firstname}</TableCell>
                  <TableCell align="center">{item.Patientrecord.Lastname}</TableCell>
                  <TableCell align="center">{item.Listofbill}</TableCell>
                  <TableCell align="center">{item.Com} บาท</TableCell>
                  <TableCell align="center">{item.Cot} บาท</TableCell>
                  <TableCell align="center">{item.Sum} บาท</TableCell>
                  <TableCell align="center">{format((new Date(item.Dateofbill)), 'dd/MM/yy')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Billlist;