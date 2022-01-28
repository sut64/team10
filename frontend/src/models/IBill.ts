import { PatientrecordInterface } from "./IPatientrecord";
export interface BillInterface{
    ID : string;
    Cot: number;
    Com: number;
    Sum: number;
    Listofbill: number;
    Dateofbill: Date; 

    PatientrecordID: number;
    Patientrecord: PatientrecordInterface;
}