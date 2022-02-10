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
import moment from 'moment';

import { AppointsInterface } from "../models/IAppoint";
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   container: {marginTop: theme.spacing(2)},
   table: { minWidth: 650},
   tableSpace: {marginTop: 20},
 })
);
 
function Appoints() {
 const classes = useStyles();
 const [appoints, setAppoints] = useState<AppointsInterface[]>([]);
 
 const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      // Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getAppoints = async () => {
    fetch(`${apiUrl}/appointments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setAppoints(res.data);
        } else {
          console.log("else");
        }
      });
  };
 
 useEffect(() => {
   getAppoints();
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
             Appointment Table
           </Typography>
         </Box>
         <Box>
           <Button
             component={RouterLink}
             to="/Appointcreate"
             variant="contained"
             color="primary"
           >
             Create Appointment
           </Button>
         </Box>
       </Box>
       <TableContainer component={Paper} className={classes.tableSpace}>
         <Table className={classes.table} aria-label="simple table">
           <TableHead>
             <TableRow>
               <TableCell align="center" width="5%">
                 No.
               </TableCell>
               <TableCell align="center" width="10%">
                  Appoint ID
               </TableCell>
               <TableCell align="center" width="20%">
                  Patient
               </TableCell>
               <TableCell align="center" width="25%">
                  Treatment
               </TableCell>
               <TableCell align="center" width="20%">
                  Doctor
               </TableCell>
               <TableCell align="center" width="5%">
                  Room number
               </TableCell>
               <TableCell align="center" width="15%">
                  Date
               </TableCell>
             </TableRow>
           </TableHead>
           <TableBody>
             {appoints.map((item: AppointsInterface) => (
               <TableRow key={item.ID}>
                 <TableCell align="center">{item.ID}</TableCell>
                 <TableCell align="center">{item.Appoint_ID}</TableCell>
                 <TableCell align="center">{item.Patientrecord.Firstname} {item.Patientrecord.Lastname}</TableCell>
                 <TableCell align="center">{item.Treatmentrecord.Treatment}</TableCell>
                 <TableCell align="center">{item.Personnel.Name}</TableCell>
                 <TableCell align="center">{item.Room_number}</TableCell>
                 <TableCell align="center">{moment(item.Date_appoint).format("DD/MM/YYYY")}</TableCell>
               </TableRow>
             ))}
           </TableBody>
         </Table>
       </TableContainer>
     </Container>
   </div>
 );
}
 
export default Appoints;
