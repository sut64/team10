import { PersonnelInterface } from "./IPersonnel";
import { PatientrecordInterface } from "./IPatientrecord";
import { DrugAllergyInterface } from "./IDrugAllergy";

export interface HistorySheetInterface {

    ID : number
    Weight : number
	Height : number
	Temperature : number
	PressureOn  : number
	PressureLow : number
    Symptom : string

	PatientrecordID : number
	Patientrecord : PatientrecordInterface;

    PersonnelID : number
	Personnel : PersonnelInterface;

	DrugAllergyID : number
	DrugAllergy : DrugAllergyInterface;

}