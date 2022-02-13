import React, { useEffect } from "react";
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
import { HistorySheetInterface } from "../models/IHistorySheet";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: { 
      marginTop: theme.spacing(2) 
    },
    table: { 
      minWidth: 1200 
    },
    tableSpace: { 
      marginTop: 20 
    },
  })
);

function HistorySheet() {
  const classes = useStyles();
  const [historysheet, setHistorySheet] = React.useState<HistorySheetInterface[]>([]);
  const apiUrl = "http://localhost:8080";
      const requestOptions = {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    };

    const getHistorySheet = async () => {
      fetch(`${apiUrl}/historysheets`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
          console.log(res.data);
          if (res.data) {
            setHistorySheet(res.data);
          } else {
            console.log("else");
          }
        });
    };

  useEffect(() => {
    getHistorySheet();
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
              History Sheet
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/HistorySheetCreate"
              variant="contained"
              color="primary"
            >
              Create History Sheet
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
              <TableCell align="center" width="4%">
                  ID
                </TableCell>
                <TableCell align="center" width="15%">
                  เจ้าหน้าที่
                </TableCell>
                <TableCell align="center" width="15%">
                  คนไข้
                </TableCell>
                <TableCell align="center" width="12%">
                  ยาที่แพ้
                </TableCell>
                <TableCell align="center" width="4%">
                  น้ำหนัก
                </TableCell>
                <TableCell align="center" width="5%">
                  ส่วนสูง
                </TableCell>
                <TableCell align="center" width="5%">
                  อุณหภูมิ
                </TableCell>
                <TableCell align="center" width="5%">
                  ความดันบน
                </TableCell>
                <TableCell align="center" width="5%">
                  ความดันล่าง
                </TableCell>
                <TableCell align="center" width="30%">
                  อาการเบื้องต้น
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {historysheet.map((history_sheets: HistorySheetInterface) => (
                <TableRow key={history_sheets.ID}>
                  <TableCell align="right"  size="medium"> {history_sheets.ID}                </TableCell>
                  <TableCell align="left"   size="medium"> {history_sheets.Personnel.Name}    </TableCell>
                  <TableCell align="left"   size="medium"> {history_sheets.Patientrecord.Firstname} {history_sheets.Patientrecord.Lastname}</TableCell>
                  <TableCell align="left"   size="medium"> {history_sheets.DrugAllergy.Name} </TableCell>
                  <TableCell align="right"  size="medium"> {history_sheets.Weight}            </TableCell>
                  <TableCell align="right"  size="medium"> {history_sheets.Height}            </TableCell>
                  <TableCell align="right"  size="medium"> {history_sheets.Temperature}       </TableCell>
                  <TableCell align="right"  size="medium"> {history_sheets.PressureOn}        </TableCell>
                  <TableCell align="right"  size="medium"> {history_sheets.PressureLow}       </TableCell>
                  <TableCell align="left"   size="medium"> {history_sheets.Symptom}           </TableCell>             
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}
export default HistorySheet;