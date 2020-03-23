package types

type GenesisState struct {
	States []string `json:"states"`
}

func NewGenesisState(whoIsRecords []string) GenesisState {
	return GenesisState{States: nil}
}

func ValidateGenesis(data GenesisState) error {
	//for _, record := range data.States {
		//if record.Owner == nil {
		//	return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		//}
		//if record.Value == "" {
		//	return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		//}
		//if record.Price == nil {
		//	return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		//}
	//}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		States: []string{},
	}
}