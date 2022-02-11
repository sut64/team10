import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

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

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบจัดการคนไข้นอก</h1>
        <h4>สมาชิกกลุ่ม</h4>
        <p>
        1.B6201067 นายปรมี สุริยะจันทร์โณ
        </p>
        <p>
        2.B6202897 นายธนวัฒน์ เหล่านคร
        </p>
        <p>
        3.B6216023 นางสาวเกษราภรณ์ เพชรนอก
        </p>
        <p>
        4.B6220297 นายนคร ศรีสรรณ์
        </p>
        <p>
        5.B6225155 นายวิทย พิลาตัน
        </p>
        <p>
        6.B6227166 นายศิริชัย ประสพผล
        </p>
      </Container>
    </div>
  );
}
export default Home;
