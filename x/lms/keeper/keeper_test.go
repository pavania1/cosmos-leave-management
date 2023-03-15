package keeper_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pavania1/cosmos-leave-management/x/lms/keeper"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
}

func (s *TestSuite) SetupTest() {
	fmt.Println("I am in setup")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(encCfg.Codec, lmsKey)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.stdntKeeper = keeper
	s.ctx = ctx
}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

//------------------------------------>> Test for AddStudent <<----------------------------------------
func (s *TestSuite) TestAddStudent() {
	tests := []struct {
		Address string
		Admin   string
		Name    string
		Id      string
	}{
		{"234", "ad", "pavani", "21"},
		{"123", "xyz", "pavs", "567"},
	}
	for _, test := range tests {
		err := s.stdntKeeper.AddStudent(s.ctx, &types.AddStudentRequest{
			Address: test.Address,
			Admin:   test.Admin,
			Name:    test.Name,
			Id:      test.Id,
		})
		s.Require().NoError(err)
	}
	Ftests := []struct {
		Address string
		Admin   string
		Name    string
		Id      string
	}{
		{"234", "ad", "pavani", "21"},
		{"123", "xyz", "pavs", "567"},
		{"456", "abcd", "pav", ""},
	}
	for _, test := range Ftests {
		err := s.stdntKeeper.CheckAddstudent(s.ctx, test.Address)
		if err == false {
			log.Println("did not add the student")
		} else {
			log.Println("Added")

		}
	}

}

//-------------------------------------->> Test for REGISTER ADMIN <<------------------------------------
func (s *TestSuite) TestRegisterAdmin() {
	tests := []struct {
		Name    string
		Address string
	}{

		{"Pavani", "234"},
		{"xyz", "567"},
	}
	for _, test := range tests {
		err := s.stdntKeeper.AdminRegister(s.ctx, &types.RegisterAdminRequest{
			Address: test.Address,
			Name:    test.Name,
		})
		s.Require().NoError(err)
	}
	Ftests := []struct {
		Name    string
		Address string
		Status  bool
	}{
		{"Pavani", "234", true},
		{"xyz", "567", true},
		{"yashhh", "123", false},
	}
	for _, test := range Ftests {
		check := s.stdntKeeper.CheckAdminregister(s.ctx, test.Address)

		s.Require().Equal(test.Status, check)

	}

}

// func (s *TestSuite) TestApplyLeave() {
// 	tests := []struct {
// 		Name    string
// 		Address string
// 		Reason string
// 		From time.Time
// 		To time.Time

// 	}{
// 		{"Pavani", "234","cold",2023-2-2,true},
// 		{"xyz", "567","fever",2023-2-2,false},
// 		{"var","123","pain",2023-2-2,true},
// 	}
// 	for _, test := range tests {
// 		err := s.stdntKeeper.ApplyLeave(s.ctx, &types.ApplyleaveRequest{
// 			Address: test.Address,
// 			Name:    test.Name,
// 			Reason: test.Address,
// 			From : test.&fromdate
// 			To: test.&todate
// 		})
// 		s.Require().NoError(err)
// 	}
//  }

//--------------------------------------------->> Test for APPLY LEAVE  <<-------------------------------
func (s *TestSuite) TestLeave() {
	dateString := "2023-02-22"
	fromdate, _ := time.Parse("2023-02-22", dateString)
	todate, _ := time.Parse("2023-02-22", "2023-02-24")

	err := s.stdntKeeper.ApplyLeave(s.ctx, &types.ApplyLeaveRequest{
		Address: "345",
		Reason:  "due to sick",
		From:    &fromdate,
		To:      &todate,
	})
	err2 := s.stdntKeeper.ApplyLeave(s.ctx, &types.ApplyLeaveRequest{
		Address: "456",
		Reason:  "dadsfadsfsdaf",
		From:    &fromdate,
		To:      &todate,
	})
	fmt.Println(err2)
	res := s.stdntKeeper.GetLeaveRqst(s.ctx, &types.GetLeaveRequest{})
	fmt.Println(res)
	s.Require().NoError(err)
}

// func (s *TestSuite) TestAcceptLeave() {
// 	err := s.stdntKeeper.AcceptLeave(s.ctx, &types.AcceptLeaveRequest{
// 		Admin:   "Pavani",
// 		LeaveId: 123,
// 		Status:  types.LeaveStatus_STATUS_ACCEPTED,
// 	})
// 	s.Require().NoError(err)
// }

//-------------------------------------->>Test fro ACCEPT LEAVE <<-------------------------------------
func (s *TestSuite) TestAcceptLeave() {
	tests := []struct {
		Admin   string
		LeaveId uint64
		Status  types.LeaveStatus
	}{
		{"pavani", 123, types.LeaveStatus_STATUS_UNDEFINED},
		{"pav", 12, types.LeaveStatus_STATUS_ACCEPTED},
		{"p", 23, types.LeaveStatus_STATUS_REJECTED},
	}
	for _, test := range tests {
		err := s.stdntKeeper.AcceptLeave(s.ctx, &types.AcceptLeaveRequest{
			Admin:   test.Admin,
			LeaveId: test.LeaveId,
			Status:  test.Status,
		})
		s.Require().NoError(err)
	}
}
func (s *TestSuite) TestGetAcceptLeave() {
	s.TestRegisterAdmin()
	s.TestAddStudent()
	s.TestLeave()
	s.TestAcceptLeave()

	res := s.stdntKeeper.GetAcceptLeaves(s.ctx, &types.GetLeaveApproveRequest{})
	fmt.Println(res)

}

// func (s *TestSuite) TestAddstudent() {
// 	except := lms.student{
// 		Id:    123,
// 		Name:  "pavani",
// 		Leave: "sick leave",
// 	}
// 	err := s.stdntKeeper.AddStudent(s.ctx, except)
// 	s.Require().NoError(err)
// 	actual, id := s.stdntKeeper.GetStudent(s.ctx, 123)
// 	s.Require().True(id)
// 	s.Require().EqualValues(except, actual)

// 	students := s.stdntKeeper.Getstudent(s.ctx)
// 	s.Require().EqualValues([]*store.Type{&except}, students)
// }
