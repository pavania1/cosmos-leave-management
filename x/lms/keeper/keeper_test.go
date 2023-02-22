package keeper_test

// import (
// 	"testing"

// 	//"github.com/cosmos/cosmos-LMS/x/lms"

// 	keeper "github.com/pavania1/cosmos-LMS/x/lms/Keeper"
// 	"github.com/stretchr/testify/suite"

// 	// "github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/x/nft"
// 	// "github.com/cosmos/cosmos-sdk/x/nft"
// 	// "github.com/pavania1/cosmos-LMS/cosmos/x/lms"
// 	// "github.com/pavania1/cosmos-LMS/cosmos/x/lms/keeper"
// 	// //"github.com/cosmos/cosmos-sdk/x/nft"
// 	// "github.com/stretchr/testify/suite"
// 	// //"google.golang.org/genproto/googleapis/actions/sdk/v2"
// )

// const (
// 	teststudentname  = "studentname"
// 	teststudentId    = "studentid"
// 	teststudentleave = "studentleave"
// )

// type TestSuite struct {
// 	suite.Suite
// 	ctx           sdk.Context
// 	address       []sdk.AccAddress
// 	studentkeeper keeper.Keeper
// 	queryClient   lms.QueryClient
// 	encCfg        lmstestutil.TestEncodingconfig
// }

// func (s *TestSuite) SetUpTest() {
// 	s.encCfg = lmstestutil.MakeTestEncodingConfig(student.AppModuleBasic{})

// 	key := sdk.NewKVStoreKey(lms.StoreKey)
// 	studentkeeper := keeper.NewKeeper(key, s.encCfg.codec)
// 	queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
// 	lms.RegisterQueryServer(queryHelper, studentkeeper)

// 	s.studentkeeper = studentkeeper
// 	s.queryClient = nft.NewQueryClient(queryHelper)
// 	s.ctx = ctx
// }
// func TestTestSuite(t *testing.T) {
// 	suite.Run(t, new(TestSuite))
// }
// func (s *TestSuite) TestAddstudent() {
// 	except := lms.student{
// 		Id:    teststudentId,
// 		Name:  teststudentname,
// 		Leave: teststudentleave,
// 	}
// 	err := s.studentkeeper.AddStudent(s.ctx, except)
// 	s.Require().NoError(err)
// 	actual, id := s.studentkeeper.GetStudent(s.ctx, teststudentId)
// 	s.Require().True(id)
// 	s.Require().EqualValues(except, actual)

// 	students := s.studentkeeper.Getstudents(s.ctx)
// 	s.Require().EqualValues([]*lms{&except}, students)
// }
// func (s *TestSuite) TestRegisterAdmin() {
// 	except := lms.Admin{
// 		Id:   testAdminId,
// 		Name: testAdminname,
// 	}
// 	err := s.studentkeeper.AdminRegister(s.ctx, except)
// 	s.Require().NoError(err)
// 	actual, id := s.studentkeeper.GetAdmin(s.ctx, testAdminId)
// 	s.Require().True(id)
// 	s.Require().EqualValues(except, actual)

// 	Admin := s.studentkeeper.GetAdmin(s.ctx)
// 	s.Require().EqualValues([]*lms{&except}, Admin)
// }
// func (s *TestSuite) TestAcceptLeave(){
// 	except:=lms.Leave{
// 		Id: testleaveid,
// 		Name: teststudentname,
// 		status: teststatusleave,
// 	}
// 	err:=s.studentKeeper.AcceptLeave(s.ctx,except)
// 	s.Require().NoError(err)
// 	actual,id:=s.LeaveStudentKeeper.GetAcceptLeave(s.ctx,testleaveid)
// 	s.Require().True(id)
// 	s.Require().Equalvalues(except,actual)

// 	leave:=s.LeaveStudentKeeper.GetAcceptLeave(s.ctx)
// 	s.Require().EqualValues([]*lms.{&except},Leave)
// }
import (
	"fmt"
	"testing"

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

func (s *TestSuite) TestAddStudent() {
	err := s.stdntKeeper.AddStudent(s.ctx, &types.AddStudentRequest{
		Address: "000",
		Admin:   "avd",
		Name:    "pavani",
		Id:      "99",
	})
	s.Require().NoError(err)

}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
func (s *TestSuite) TestRegisterAdmin() {
	err := s.stdntKeeper.AdminRegister(s.ctx, &types.RegisterAdminRequest{
		Address: "567",
		Name:    "xyz",
	})
	s.Require().NoError(err)
}
