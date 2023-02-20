package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group/keeper"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/actions/sdk/v2"
)

const (
	teststudentname  = "studentname"
	teststudentId    = "studentid"
	teststudentleave = "studentleave"
)

type TestSuite struct {
	suite.Suite
	ctx           sdk.Context
	address       []sdk.AccAddress
	studentkeeper keeper.Keeper
	//encCfg        studetestutil.TestEncodingconfig
}

func (t *Addstudent) TestAddstudent() {

}
