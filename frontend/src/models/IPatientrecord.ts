import { BloodTypeInterface } from "./IBloodtype";
import { GenderInterface } from "./IGender";
import { PersonnelInterface } from "./IPersonnel";
import { PrenamesInterface } from "./IPrename";
import { ProvincesInterface } from "./IProvince";

export interface PatientrecordInterface {
  ID: string,
  PrenameID: number,
  Prename: PrenamesInterface,
  Firstname: string,
  Lastname: string,
  GenderID: number,
  Gender: GenderInterface,
  Idcardnumber: string,
  Age: number,
  Birthday: Date,
  BloodTypeID: number,
  BloodType: BloodTypeInterface,
  Phonenumber: number,
  Email: string,
  Home: string,
  ProvinceID: number,
  Province: ProvincesInterface,
  Emergencyname: string,
  Emergencyphone: number,
  Timestamp: Date,
  PersonnelID: number,
  Personnel: PersonnelInterface,
}
