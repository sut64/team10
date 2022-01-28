import { BloodTypeInterface } from "./IBloodtype";
import { GenderInterface } from "./IGender";
import { JobTitleInterface } from "./IJobtitle";

export interface PersonnelInterface {

    ID: number,
	Name:       string;
	Personalid:  string;   
	BirthDay:    Date;
	Tel    :     string;
	Address :    string;
	Salary :     number;

	GenderID   : number;
	Gender     : GenderInterface;

	BloodTypeID :number;
	BloodType :  BloodTypeInterface;
    
	JobTitleID :number;
	JobTitle   : JobTitleInterface;
   }
   