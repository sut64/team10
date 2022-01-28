import { DiseaseInterface } from "./IDisease";
import { MedicineInterface } from "./IMedicine";
import { PatientrecordInterface } from "./IPatientrecord"
import { PersonnelInterface } from "./IPersonnel";
export interface TreatmentrecordInterface {
    ID: number,
    Treatment: string,
    Temperature: string,
    Date: Date,
    DiseaseID: number,
    Disease: DiseaseInterface,
    MedicineID: number,
    Medicine: MedicineInterface,
    PatientrecordID: number,
    Patient: PatientrecordInterface,
    PersonnelID: number,
    Personnel: PersonnelInterface,
}