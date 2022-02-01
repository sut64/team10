import { PatientrecordInterface } from "./IPatientrecord";
import { TreatmentrecordInterface } from "./ITreatmentrecord";
import { PersonnelInterface } from "./IPersonnel";

export interface AppointsInterface {
    ID: number,
    Appoint_ID: string,
    Room_number: number,
    Date_appoint: Date | null,

    PersonnelID: number,
    Personnel: PersonnelInterface,

    PatientrecordID: number,
    Patientrecord: PatientrecordInterface,

    TreatmentrecordID: number,
    Treatmentrecord: TreatmentrecordInterface,
   }
   
