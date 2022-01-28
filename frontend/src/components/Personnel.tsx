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
import { PersonnelInterface } from "../models/IPersonnel";
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

function Personnel() {
  const classes = useStyles();
  const [personnel, setPersonnel] = useState<PersonnelInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const getPersonnel = async () => {
    fetch(`${apiUrl}/personnel`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setPersonnel(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getPersonnel();
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
              ระบบการข้อมูลบุคลลากร
            </Typography>
            <Box>
            <Button
              component={RouterLink}
              to="/personnel/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  รหัส
                </TableCell>
                <TableCell align="center" width="28%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="15%">
                  อาชีพ
                </TableCell>
                <TableCell align="center" width="5%">
                  เพศ
                </TableCell>
                <TableCell align="center" width="5%">
                  กรุ๊ปเลือด
                </TableCell>
                <TableCell align="center" width="17%">
                  เบอร์โทร
                </TableCell>
                <TableCell align="center" width="25%">
                  วันเกิด
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {personnel.map((item: PersonnelInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Name}</TableCell>
                  <TableCell align="center">{item.JobTitle.Job}</TableCell>
                  <TableCell align="center">{item.Gender.Genders}</TableCell>
                  <TableCell align="center">{item.BloodType.BloodType}</TableCell>
                  <TableCell align="center">{item.Tel}</TableCell>
                  <TableCell align="center">{format((new Date(item.BirthDay)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}

            </TableBody>
          </Table> 
          
        </TableContainer>
      </Container>
    </div>
  );
}

export default Personnel;
