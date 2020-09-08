package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation/d2parser"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// SkillDetails has all of the SkillRecords
//nolint:gochecknoglobals // Currently global by design, only written once
var SkillDetails map[int]*SkillRecord

// SkillRecord is a row from the skills.txt file. Here are two resources for more info on each field
// [https://d2mods.info/forum/viewtopic.php?t=41556, https://d2mods.info/forum/kb/viewarticle?a=246]
type SkillRecord struct {
	Skill             string
	Charclass         string
	Skilldesc         string
	Prgcalc1          d2calculation.Calculation
	Prgcalc2          d2calculation.Calculation
	Prgcalc3          d2calculation.Calculation
	Srvmissile        string
	Srvmissilea       string
	Srvmissileb       string
	Srvmissilec       string
	Srvoverlay        string
	Aurastate         string
	Auratargetstate   string
	Auralencalc       d2calculation.Calculation
	Aurarangecalc     d2calculation.Calculation
	Aurastat1         string
	Aurastatcalc1     d2calculation.Calculation
	Aurastat2         string
	Aurastatcalc2     d2calculation.Calculation
	Aurastat3         string
	Aurastatcalc3     d2calculation.Calculation
	Aurastat4         string
	Aurastatcalc4     d2calculation.Calculation
	Aurastat5         string
	Aurastatcalc5     d2calculation.Calculation
	Aurastat6         string
	Aurastatcalc6     d2calculation.Calculation
	Auraevent1        string
	Auraevent2        string
	Auraevent3        string
	Auratgtevent      string
	Auratgteventfunc  string
	Passivestate      string
	Passiveitype      string
	Passivestat1      string
	Passivecalc1      d2calculation.Calculation
	Passivestat2      string
	Passivecalc2      d2calculation.Calculation
	Passivestat3      string
	Passivecalc3      d2calculation.Calculation
	Passivestat4      string
	Passivecalc4      d2calculation.Calculation
	Passivestat5      string
	Passivecalc5      d2calculation.Calculation
	Passiveevent      string
	Passiveeventfunc  string
	Summon            string
	Pettype           string
	Petmax            d2calculation.Calculation
	Summode           string
	Sumskill1         string
	Sumsk1calc        d2calculation.Calculation
	Sumskill2         string
	Sumsk2calc        d2calculation.Calculation
	Sumskill3         string
	Sumsk3calc        d2calculation.Calculation
	Sumskill4         string
	Sumsk4calc        d2calculation.Calculation
	Sumskill5         string
	Sumsk5calc        d2calculation.Calculation
	Sumoverlay        string
	Stsound           string
	Stsoundclass      string
	Dosound           string
	DosoundA          string
	DosoundB          string
	Tgtoverlay        string
	Tgtsound          string
	Prgoverlay        string
	Prgsound          string
	Castoverlay       string
	Cltoverlaya       string
	Cltoverlayb       string
	Cltmissile        string
	Cltmissilea       string
	Cltmissileb       string
	Cltmissilec       string
	Cltmissiled       string
	Cltcalc1          d2calculation.Calculation
	Cltcalc2          d2calculation.Calculation
	Cltcalc3          d2calculation.Calculation
	Range             string
	Itypea1           string
	Itypea2           string
	Itypea3           string
	Etypea1           string
	Etypea2           string
	Itypeb1           string
	Itypeb2           string
	Itypeb3           string
	Etypeb1           string
	Etypeb2           string
	Anim              string
	Seqtrans          string
	Monanim           string
	ItemCastSound     string
	ItemCastOverlay   string
	Skpoints          d2calculation.Calculation
	Reqskill1         string
	Reqskill2         string
	Reqskill3         string
	State1            string
	State2            string
	State3            string
	Perdelay          d2calculation.Calculation
	Calc1             d2calculation.Calculation
	Calc2             d2calculation.Calculation
	Calc3             d2calculation.Calculation
	Calc4             d2calculation.Calculation
	ToHitCalc         d2calculation.Calculation
	DmgSymPerCalc     d2calculation.Calculation
	EType             string
	EDmgSymPerCalc    d2calculation.Calculation
	ELenSymPerCalc    d2calculation.Calculation
	ID                int
	Srvstfunc         int
	Srvdofunc         int
	Srvprgfunc1       int
	Srvprgfunc2       int
	Srvprgfunc3       int
	Prgdam            int
	Aurafilter        int
	Auraeventfunc1    int
	Auraeventfunc2    int
	Auraeventfunc3    int
	Sumumod           int
	Cltstfunc         int
	Cltdofunc         int
	Cltprgfunc1       int
	Cltprgfunc2       int
	Cltprgfunc3       int
	Attackrank        int
	Weapsel           int
	Seqnum            int
	Seqinput          int
	LineOfSight       int
	SelectProc        int
	ItemEffect        int
	ItemCltEffect     int
	ItemTgtDo         int
	ItemTarget        int
	Reqlevel          int
	Maxlvl            int
	Reqstr            int
	Reqdex            int
	Reqint            int
	Reqvit            int
	Restrict          int
	Delay             int
	Checkfunc         int
	Startmana         int
	Minmana           int
	Manashift         int
	Mana              int
	Lvlmana           int
	Param1            int
	Param2            int
	Param3            int
	Param4            int
	Param5            int
	Param6            int
	Param7            int
	Param8            int
	ToHit             int
	LevToHit          int
	ResultFlags       int
	HitFlags          int
	HitClass          int
	HitShift          int
	SrcDam            int
	MinDam            int
	MinLevDam1        int
	MinLevDam2        int
	MinLevDam3        int
	MinLevDam4        int
	MinLevDam5        int
	MaxDam            int
	MaxLevDam1        int
	MaxLevDam2        int
	MaxLevDam3        int
	MaxLevDam4        int
	MaxLevDam5        int
	EMin              int
	EMinLev1          int
	EMinLev2          int
	EMinLev3          int
	EMinLev4          int
	EMinLev5          int
	EMax              int
	EMaxLev1          int
	EMaxLev2          int
	EMaxLev3          int
	EMaxLev4          int
	EMaxLev5          int
	ELen              int
	ELevLen1          int
	ELevLen2          int
	ELevLen3          int
	Aitype            int
	Aibonus           int
	CostMult          int
	CostAdd           int
	Prgstack          bool
	Decquant          bool
	Lob               bool
	Stsuccessonly     bool
	Stsounddelay      bool
	Weaponsnd         bool
	Warp              bool
	Immediate         bool
	Enhanceable       bool
	Noammo            bool
	Durability        bool
	UseAttackRate     bool
	TargetableOnly    bool
	SearchEnemyXY     bool
	SearchEnemyNear   bool
	SearchOpenXY      bool
	TargetCorpse      bool
	TargetPet         bool
	TargetAlly        bool
	TargetItem        bool
	AttackNoMana      bool
	TgtPlaceCheck     bool
	ItemCheckStart    bool
	ItemCltCheckStart bool
	Leftskill         bool
	Repeat            bool
	Nocostinstate     bool
	Usemanaondo       bool
	Interrupt         bool
	InTown            bool
	Aura              bool
	Periodic          bool
	Finishing         bool
	Passive           bool
	Progressive       bool
	General           bool
	Scroll            bool
	InGame            bool
	Kick              bool
}

// LoadSkills loads skills.txt file contents into a skill record map
//nolint:funlen // Makes no sense to split
// LoadCharStats loads charstats.txt file contents into map[d2enum.Hero]*CharStatsRecord
func LoadSkills(file []byte) {
	SkillDetails = make(map[int]*SkillRecord)

	parser := d2parser.New()

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		name := d.String("skill")
		parser.SetCurrentReference("skill", name)

		record := &SkillRecord{
			Skill:             d.String("skill"),
			ID:                d.Number("Id"),
			Charclass:         d.String("charclass"),
			Skilldesc:         d.String("skilldesc"),
			Srvstfunc:         d.Number("srvstfunc"),
			Srvdofunc:         d.Number("srvdofunc"),
			Prgstack:          d.Bool("prgstack"),
			Srvprgfunc1:       d.Number("srvprgfunc1"),
			Srvprgfunc2:       d.Number("srvprgfunc2"),
			Srvprgfunc3:       d.Number("srvprgfunc3"),
			Prgcalc1:          parser.Parse(d.String("prgcalc1")),
			Prgcalc2:          parser.Parse(d.String("prgcalc2")),
			Prgcalc3:          parser.Parse(d.String("prgcalc3")),
			Prgdam:            d.Number("prgdam"),
			Srvmissile:        d.String("srvmissile"),
			Decquant:          d.Bool("decquant"),
			Lob:               d.Bool("lob"),
			Srvmissilea:       d.String("srvmissilea"),
			Srvmissileb:       d.String("srvmissileb"),
			Srvmissilec:       d.String("srvmissilec"),
			Srvoverlay:        d.String("srvoverlay"),
			Aurafilter:        d.Number("aurafilter"),
			Aurastate:         d.String("aurastate"),
			Auratargetstate:   d.String("auratargetstate"),
			Auralencalc:       parser.Parse(d.String("auralencalc")),
			Aurarangecalc:     parser.Parse(d.String("aurarangecalc")),
			Aurastat1:         d.String("aurastat1"),
			Aurastatcalc1:     parser.Parse(d.String("aurastatcalc1")),
			Aurastat2:         d.String("aurastat2"),
			Aurastatcalc2:     parser.Parse(d.String("aurastatcalc2")),
			Aurastat3:         d.String("aurastat3"),
			Aurastatcalc3:     parser.Parse(d.String("aurastatcalc3")),
			Aurastat4:         d.String("aurastat4"),
			Aurastatcalc4:     parser.Parse(d.String("aurastatcalc4")),
			Aurastat5:         d.String("aurastat5"),
			Aurastatcalc5:     parser.Parse(d.String("aurastatcalc5")),
			Aurastat6:         d.String("aurastat6"),
			Aurastatcalc6:     parser.Parse(d.String("aurastatcalc6")),
			Auraevent1:        d.String("auraevent1"),
			Auraeventfunc1:    d.Number("auraeventfunc1"),
			Auraevent2:        d.String("auraevent2"),
			Auraeventfunc2:    d.Number("auraeventfunc2"),
			Auraevent3:        d.String("auraevent3"),
			Auraeventfunc3:    d.Number("auraeventfunc3"),
			Auratgtevent:      d.String("auratgtevent"),
			Auratgteventfunc:  d.String("auratgteventfunc"),
			Passivestate:      d.String("passivestate"),
			Passiveitype:      d.String("passiveitype"),
			Passivestat1:      d.String("passivestat1"),
			Passivecalc1:      parser.Parse(d.String("passivecalc1")),
			Passivestat2:      d.String("passivestat2"),
			Passivecalc2:      parser.Parse(d.String("passivecalc2")),
			Passivestat3:      d.String("passivestat3"),
			Passivecalc3:      parser.Parse(d.String("passivecalc3")),
			Passivestat4:      d.String("passivestat4"),
			Passivecalc4:      parser.Parse(d.String("passivecalc4")),
			Passivestat5:      d.String("passivestat5"),
			Passivecalc5:      parser.Parse(d.String("passivecalc5")),
			Passiveevent:      d.String("passiveevent"),
			Passiveeventfunc:  d.String("passiveeventfunc"),
			Summon:            d.String("summon"),
			Pettype:           d.String("pettype"),
			Petmax:            parser.Parse(d.String("petmax")),
			Summode:           d.String("summode"),
			Sumskill1:         d.String("sumskill1"),
			Sumsk1calc:        parser.Parse(d.String("sumsk1calc")),
			Sumskill2:         d.String("sumskill2"),
			Sumsk2calc:        parser.Parse(d.String("sumsk2calc")),
			Sumskill3:         d.String("sumskill3"),
			Sumsk3calc:        parser.Parse(d.String("sumsk3calc")),
			Sumskill4:         d.String("sumskill4"),
			Sumsk4calc:        parser.Parse(d.String("sumsk4calc")),
			Sumskill5:         d.String("sumskill5"),
			Sumsk5calc:        parser.Parse(d.String("sumsk5calc")),
			Sumumod:           d.Number("sumumod"),
			Sumoverlay:        d.String("sumoverlay"),
			Stsuccessonly:     d.Bool("stsuccessonly"),
			Stsound:           d.String("stsound"),
			Stsoundclass:      d.String("stsoundclass"),
			Stsounddelay:      d.Bool("stsounddelay"),
			Weaponsnd:         d.Bool("weaponsnd"),
			Dosound:           d.String("dosound"),
			DosoundA:          d.String("dosound a"),
			DosoundB:          d.String("dosound b"),
			Tgtoverlay:        d.String("tgtoverlay"),
			Tgtsound:          d.String("tgtsound"),
			Prgoverlay:        d.String("prgoverlay"),
			Prgsound:          d.String("prgsound"),
			Castoverlay:       d.String("castoverlay"),
			Cltoverlaya:       d.String("cltoverlaya"),
			Cltoverlayb:       d.String("cltoverlayb"),
			Cltstfunc:         d.Number("cltstfunc"),
			Cltdofunc:         d.Number("cltdofunc"),
			Cltprgfunc1:       d.Number("cltprgfunc1"),
			Cltprgfunc2:       d.Number("cltprgfunc2"),
			Cltprgfunc3:       d.Number("cltprgfunc3"),
			Cltmissile:        d.String("cltmissile"),
			Cltmissilea:       d.String("cltmissilea"),
			Cltmissileb:       d.String("cltmissileb"),
			Cltmissilec:       d.String("cltmissilec"),
			Cltmissiled:       d.String("cltmissiled"),
			Cltcalc1:          parser.Parse(d.String("cltcalc1")),
			Cltcalc2:          parser.Parse(d.String("cltcalc2")),
			Cltcalc3:          parser.Parse(d.String("cltcalc3")),
			Warp:              d.Bool("warp"),
			Immediate:         d.Bool("immediate"),
			Enhanceable:       d.Bool("enhanceable"),
			Attackrank:        d.Number("attackrank"),
			Noammo:            d.Bool("noammo"),
			Range:             d.String("range"),
			Weapsel:           d.Number("weapsel"),
			Itypea1:           d.String("itypea1"),
			Itypea2:           d.String("itypea2"),
			Itypea3:           d.String("itypea3"),
			Etypea1:           d.String("etypea1"),
			Etypea2:           d.String("etypea2"),
			Itypeb1:           d.String("itypeb1"),
			Itypeb2:           d.String("itypeb2"),
			Itypeb3:           d.String("itypeb3"),
			Etypeb1:           d.String("etypeb1"),
			Etypeb2:           d.String("etypeb2"),
			Anim:              d.String("anim"),
			Seqtrans:          d.String("seqtrans"),
			Monanim:           d.String("monanim"),
			Seqnum:            d.Number("seqnum"),
			Seqinput:          d.Number("seqinput"),
			Durability:        d.Bool("durability"),
			UseAttackRate:     d.Bool("UseAttackRate"),
			LineOfSight:       d.Number("LineOfSight"),
			TargetableOnly:    d.Bool("TargetableOnly"),
			SearchEnemyXY:     d.Bool("SearchEnemyXY"),
			SearchEnemyNear:   d.Bool("SearchEnemyNear"),
			SearchOpenXY:      d.Bool("SearchOpenXY"),
			SelectProc:        d.Number("SelectProc"),
			TargetCorpse:      d.Bool("TargetCorpse"),
			TargetPet:         d.Bool("TargetPet"),
			TargetAlly:        d.Bool("TargetAlly"),
			TargetItem:        d.Bool("TargetItem"),
			AttackNoMana:      d.Bool("AttackNoMana"),
			TgtPlaceCheck:     d.Bool("TgtPlaceCheck"),
			ItemEffect:        d.Number("ItemEffect"),
			ItemCltEffect:     d.Number("ItemCltEffect"),
			ItemTgtDo:         d.Number("ItemTgtDo"),
			ItemTarget:        d.Number("ItemTarget"),
			ItemCheckStart:    d.Bool("ItemCheckStart"),
			ItemCltCheckStart: d.Bool("ItemCltCheckStart"),
			ItemCastSound:     d.String("ItemCastSound"),
			ItemCastOverlay:   d.String("ItemCastOverlay"),
			Skpoints:          parser.Parse(d.String("skpoints")),
			Reqlevel:          d.Number("reqlevel"),
			Maxlvl:            d.Number("maxlvl"),
			Reqstr:            d.Number("reqstr"),
			Reqdex:            d.Number("reqdex"),
			Reqint:            d.Number("reqint"),
			Reqvit:            d.Number("reqvit"),
			Reqskill1:         d.String("reqskill1"),
			Reqskill2:         d.String("reqskill2"),
			Reqskill3:         d.String("reqskill3"),
			Restrict:          d.Number("restrict"),
			State1:            d.String("State1"),
			State2:            d.String("State2"),
			State3:            d.String("State3"),
			Delay:             d.Number("delay"),
			Leftskill:         d.Bool("leftskill"),
			Repeat:            d.Bool("repeat"),
			Checkfunc:         d.Number("checkfunc"),
			Nocostinstate:     d.Bool("nocostinstate"),
			Usemanaondo:       d.Bool("usemanaondo"),
			Startmana:         d.Number("startmana"),
			Minmana:           d.Number("minmana"),
			Manashift:         d.Number("manashift"),
			Mana:              d.Number("mana"),
			Lvlmana:           d.Number("lvlmana"),
			Interrupt:         d.Bool("interrupt"),
			InTown:            d.Bool("InTown"),
			Aura:              d.Bool("aura"),
			Periodic:          d.Bool("periodic"),
			Perdelay:          parser.Parse(d.String("perdelay")),
			Finishing:         d.Bool("finishing"),
			Passive:           d.Bool("passive"),
			Progressive:       d.Bool("progressive"),
			General:           d.Bool("general"),
			Scroll:            d.Bool("scroll"),
			Calc1:             parser.Parse(d.String("calc1")),
			Calc2:             parser.Parse(d.String("calc2")),
			Calc3:             parser.Parse(d.String("calc3")),
			Calc4:             parser.Parse(d.String("calc4")),
			Param1:            d.Number("Param1"),
			Param2:            d.Number("Param2"),
			Param3:            d.Number("Param3"),
			Param4:            d.Number("Param4"),
			Param5:            d.Number("Param5"),
			Param6:            d.Number("Param6"),
			Param7:            d.Number("Param7"),
			Param8:            d.Number("Param8"),
			InGame:            d.Bool("InGame"),
			ToHit:             d.Number("ToHit"),
			LevToHit:          d.Number("LevToHit"),
			ToHitCalc:         parser.Parse(d.String("ToHitCalc")),
			ResultFlags:       d.Number("ResultFlags"),
			HitFlags:          d.Number("HitFlags"),
			HitClass:          d.Number("HitClass"),
			Kick:              d.Bool("Kick"),
			HitShift:          d.Number("HitShift"),
			SrcDam:            d.Number("SrcDam"),
			MinDam:            d.Number("MinDam"),
			MinLevDam1:        d.Number("MinLevDam1"),
			MinLevDam2:        d.Number("MinLevDam2"),
			MinLevDam3:        d.Number("MinLevDam3"),
			MinLevDam4:        d.Number("MinLevDam4"),
			MinLevDam5:        d.Number("MinLevDam5"),
			MaxDam:            d.Number("MaxDam"),
			MaxLevDam1:        d.Number("MaxLevDam1"),
			MaxLevDam2:        d.Number("MaxLevDam2"),
			MaxLevDam3:        d.Number("MaxLevDam3"),
			MaxLevDam4:        d.Number("MaxLevDam4"),
			MaxLevDam5:        d.Number("MaxLevDam5"),
			DmgSymPerCalc:     parser.Parse(d.String("DmgSymPerCalc")),
			EType:             d.String("EType"),
			EMin:              d.Number("EMin"),
			EMinLev1:          d.Number("EMinLev1"),
			EMinLev2:          d.Number("EMinLev2"),
			EMinLev3:          d.Number("EMinLev3"),
			EMinLev4:          d.Number("EMinLev4"),
			EMinLev5:          d.Number("EMinLev5"),
			EMax:              d.Number("EMax"),
			EMaxLev1:          d.Number("EMaxLev1"),
			EMaxLev2:          d.Number("EMaxLev2"),
			EMaxLev3:          d.Number("EMaxLev3"),
			EMaxLev4:          d.Number("EMaxLev4"),
			EMaxLev5:          d.Number("EMaxLev5"),
			EDmgSymPerCalc:    parser.Parse(d.String("EDmgSymPerCalc")),
			ELen:              d.Number("ELen"),
			ELevLen1:          d.Number("ELevLen1"),
			ELevLen2:          d.Number("ELevLen2"),
			ELevLen3:          d.Number("ELevLen3"),
			ELenSymPerCalc:    parser.Parse(d.String("ELenSymPerCalc")),
			Aitype:            d.Number("aitype"),
			Aibonus:           d.Number("aibonus"),
			CostMult:          d.Number("cost mult"),
			CostAdd:           d.Number("cost add"),
		}
		SkillDetails[record.ID] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Skill records", len(SkillDetails))
}
