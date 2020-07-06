package main

// Study is the structure of the Study instance
type Study struct {
	Name        string `json:"name"`
	Mrn         string `json:"mrn"`
	Accession   string `json:"accession"`
	StudyDate   string `json:"studyDate"`
	Modality    string `json:"modality"`
	Description string `json:"description"`
	Selected    bool   `json:"selected"`
}

// Studies Represents a List of MOC Studies
var Studies = []Study{
	{
		Name:        "Carmelia Tunnell",
		Mrn:         "8277166112",
		Accession:   "4406894353",
		StudyDate:   "01/11/2020",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    true,
	},
	{
		Name:        "Karen Lindgren",
		Mrn:         "6715972370",
		Accession:   "5581442436",
		StudyDate:   "05/11/2008",
		Modality:    "SM\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Cherilyn Fuhrman",
		Mrn:         "2807966769",
		Accession:   "3129884101",
		StudyDate:   "02/14/2018",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Janna Linebarger",
		Mrn:         "9350719467",
		Accession:   "4623607356",
		StudyDate:   "07/16/2002",
		Modality:    "CT\\PT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Nakita Cargle",
		Mrn:         "9232361637",
		Accession:   "7715405825",
		StudyDate:   "02/01/2011",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Allene Treadway",
		Mrn:         "2954695563",
		Accession:   "2151889651",
		StudyDate:   "01/20/2007",
		Modality:    "SM\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Kari Medley",
		Mrn:         "4140544897",
		Accession:   "1921754564",
		StudyDate:   "07/01/2013",
		Modality:    "CT\\PT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Mario Rusk",
		Mrn:         "4319709897",
		Accession:   "4059571607",
		StudyDate:   "10/23/2002",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Lula Blanks",
		Mrn:         "2063533309",
		Accession:   "3038157963",
		StudyDate:   "04/25/2011",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Sommer Bedard",
		Mrn:         "7231442795",
		Accession:   "4122383012",
		StudyDate:   "12/02/2006",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Tod Herbert",
		Mrn:         "210944890",
		Accession:   "3849111281",
		StudyDate:   "09/16/2012",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
	{
		Name:        "Trisha Cambell",
		Mrn:         "3418785970",
		Accession:   "1901727335",
		StudyDate:   "02/26/2011",
		Modality:    "CT\\SR",
		Description: "None",
		Selected:    false,
	},
}
