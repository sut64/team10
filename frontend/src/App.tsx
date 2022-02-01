import React from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Switch, Route, Link} from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";

import HomeIcon from "@material-ui/icons/Home";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
import LibraryBookIcon from "@material-ui/icons/LibraryBooks";
import BookIcon from "@material-ui/icons/Book";
import AssignmentIcon from '@material-ui/icons/Assignment';
import AccessAlarmsIcon from '@material-ui/icons/AccessAlarms';

import Home from "./components/Home";
import CreatePersonnel from "./components/CreatePersonnel";
import Personnel from "./components/Personnel";
import PatientrecordCreate from "./components/PatientrecordCreate";
import Patientrecords from "./components/Patientrecords";
import CreateBill from "./components/CreateBill";
import Bill from "./components/Bill";
import HistorySheets from "./components/HistorySheets";
import CreateHistorySheet from "./components/CreateHistorySheet";
import Appoints from "./components/Appoints";
import AppointCreate from "./components/AppointCreate";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
  })
);

export default function MiniDrawer() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const menu = [
    { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },
    { name: "ข้อมูลบุคคลากร", icon: <AccountCircleIcon />, path: "/personnel" },
    { name: "ลงทะเบียนคนไข้", icon: <LibraryBookIcon />, path: "/patientrecord/create" },
    { name: "รายการบิล", icon: <BookIcon />,path: "/Bill"},
    { name: "ซักประวัติ", icon: <AssignmentIcon />,path: "/CreateHistorySheet"},
    { name: "การนัดคนไข้", icon: <AccessAlarmsIcon />,path: "/AppointTable"},
  ];

  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        <AppBar
          position="fixed"
          className={clsx(classes.appBar, {
            [classes.appBarShift]: open,
          })}
        >
          <Toolbar>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              onClick={handleDrawerOpen}
              edge="start"
              className={clsx(classes.menuButton, {
                [classes.hide]: open,
              })}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" noWrap>
                ระบบจัดการคนไข้นอก
            </Typography>
          </Toolbar>
        </AppBar>
        <Drawer
          variant="permanent"
          className={clsx(classes.drawer, {
            [classes.drawerOpen]: open,
            [classes.drawerClose]: !open,
          })}
          classes={{
            paper: clsx({
              [classes.drawerOpen]: open,
              [classes.drawerClose]: !open,
            }),
          }}
        >
          <div className={classes.toolbar}>
            <IconButton onClick={handleDrawerClose}>
              {theme.direction === "rtl" ? (
                <ChevronRightIcon />
              ) : (
                <ChevronLeftIcon />
              )}
            </IconButton>
          </div>
          <Divider />
          <List>
            {menu.map((item, index) => (
              <Link to={item.path} key={item.name}>
                <ListItem button>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText primary={item.name} />
                </ListItem>
              </Link>
            ))}
          </List>
        </Drawer>
        <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/personnel" component={Personnel} />
              <Route exact path="/CreateBill" component={CreateBill}/>
              <Route exact path="/Bill" component={Bill}/>
              <Route exact path="/HistorySheets" component={HistorySheets}/>
              <Route exact path="/CreateHistorySheet" component={CreateHistorySheet}/>
              <Route
                exact
                path="/personnel/create"
                component={CreatePersonnel} 
              />
              <Route exact path="/patientrecord/create" component={PatientrecordCreate} />
              <Route exact path="/patientrecords" component={Patientrecords} />
	      <Route exact path="/AppointTable" component={Appoints} />
              <Route exact path="/Appointcreate" component={AppointCreate} />		              

              <Route component={Home} path="/" />
              
            </Switch>
          </div>
        </main>
      </Router>
    </div>
  );


}
