import { DiseaseInterface } from "./IDisease";
import { MedicineInterface } from "./IMedicine";
import { PatientrecordInterface } from "./IPatientrecord"
import { PersonnelInterface } from "./IPersonnel";
export interface TreatmentrecordInterface {
    ID: number,
    Treatment: string,
    Temperature: number,
    DiseaseID: number,
    Disease: DiseaseInterface,
    MedicineID: number,
    Medicine: MedicineInterface,
    PatientrecordID: number,
    Patientrecord: PatientrecordInterface,
    PersonnelID: number,
    Personnel: PersonnelInterface,
    RecordDate: Date | null,
}