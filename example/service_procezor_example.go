package example

import (
	legalios "github.com/mzdyhrave/legaliosgo"
	procezor "github.com/mzdyhrave/procezorgo"
	"github.com/mzdyhrave/procezorgo/internal/types"
)

const TEST_VERSION int32 = 100

type TestConceptSpec struct {
	procezor.ConceptSpec
}

func NewTestConceptSpec(code int32) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptSpec(code)}
}

func NewTestConceptPathSpec(code int32, path []procezor.ArticleCode) TestConceptSpec {
	return TestConceptSpec{ procezor.NewConceptPathSpec(code, path)}
}

func NewTestConceptFuncSpec(code int32, resultFunc procezor.ResultFunc) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptFuncSpec(code, resultFunc)}
}

func NewTestConceptPathFuncSpec(code int32, path []procezor.ArticleCode, resultFunc procezor.ResultFunc) TestConceptSpec {
	return TestConceptSpec{procezor.NewConceptPathFuncSpec(code, path, resultFunc)}
}

type TestTermTarget struct {
	procezor.TermTarget
}

func (t TestTermTarget) ArticleDescr() string {
	return TestArticleConst(t.Article().Value()).String()
}

func (t TestTermTarget) ConceptDescr() string {
	return TestConceptConst(t.Concept().Value()).String()
}

func NewTestTermTarget(mont types.MonthCode, cont types.ContractCode, post types.PositionCode, vars types.VariantCode, code types.ArticleCode, role types.ConceptCode) TestTermTarget {
	return TestTermTarget{ TermTarget: procezor.NewTermTarget(mont, cont, post, vars, code, role) }
}

type TestTermResult struct {
	procezor.TermResult
}

func (t TestTermResult) ArticleDescr() string {
	return TestArticleConst(t.Article().Value()).String()
}

func (t TestTermResult) ConceptDescr() string {
	return TestConceptConst(t.Concept().Value()).String()
}


func NewTestTermResult(target procezor.ITermTarget, value int32, basis int32, descr string) TestTermResult {
	return TestTermResult{ TermResult: procezor.NewTermResult(target, value, basis, descr) }
}

type TestArticleFactory struct {
	procezor.IArticleSpecFactory
}

func NewTestArticleFactory() procezor.IArticleSpecFactory{
	providersConfig := []procezor.ProviderRecord {
		procezor.NewProviderRecord(ARTICLE_TIMESHT_WORKING.Id(), CONCEPT_TIMESHT_WORKING.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_SALARY.Id(), CONCEPT_AMOUNT_BASIS.Id(),
			[]int32 {
				ARTICLE_INCOME_GROSS.Id(),
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BONUS.Id(), CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_INCOME_GROSS.Id(),
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_PAYMENT_BARTER.Id(), CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_HEALTH_INSBASE.Id(),
				ARTICLE_SOCIAL_INSBASE.Id(),
				ARTICLE_TAXING_ADVBASE.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_ALLOWCE_HOFFICE.Id(), CONCEPT_AMOUNT_FIXED.Id(),
			[]int32 {
				ARTICLE_INCOME_NETTO.Id(),
			}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSBASE.Id(), CONCEPT_HEALTH_INSBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSBASE.Id(), CONCEPT_SOCIAL_INSBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_HEALTH_INSPAYM.Id(), CONCEPT_HEALTH_INSPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_SOCIAL_INSPAYM.Id(), CONCEPT_SOCIAL_INSPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVBASE.Id(), CONCEPT_TAXING_ADVBASE.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_TAXING_ADVPAYM.Id(), CONCEPT_TAXING_ADVPAYM.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_GROSS.Id(), CONCEPT_INCOME_GROSS.Id(),
			[]int32{}),
		procezor.NewProviderRecord(ARTICLE_INCOME_NETTO.Id(), CONCEPT_INCOME_NETTO.Id(),
			[]int32{}),
	}
	return TestArticleFactory{procezor.NewArticleSpecFactoryWithRecords(providersConfig)}
}

type TestConceptFactory struct {
	procezor.IConceptSpecFactory
}

func NewTestConceptFactory() procezor.IConceptSpecFactory {
	return TestConceptFactory{procezor.NewConceptSpecFactoryWithProviders(
		map[int32]procezor.IConceptSpecProvider {
			CONCEPT_TIMESHT_WORKING.Id(): NewTimeshtWorkingConProv(),
			CONCEPT_AMOUNT_BASIS.Id(): NewAmountBasisConProv(),
			CONCEPT_AMOUNT_FIXED.Id(): NewAmountFixedConProv(),
			CONCEPT_HEALTH_INSBASE.Id(): NewHealthInsbaseConProv(),
			CONCEPT_SOCIAL_INSBASE.Id(): NewSocialInsbaseConProv(),
			CONCEPT_HEALTH_INSPAYM.Id(): NewHealthInspaymConProv(),
			CONCEPT_SOCIAL_INSPAYM.Id(): NewSocialInspaymConProv(),
			CONCEPT_TAXING_ADVBASE.Id(): NewTaxingAdvbaseConProv(),
			CONCEPT_TAXING_ADVPAYM.Id(): NewTaxingAdvpaymConProv(),
			CONCEPT_INCOME_GROSS.Id(): NewIncomeGrossConProv(),
			CONCEPT_INCOME_NETTO.Id(): NewIncomeNettoConProv(),
		}),
	}
}

type TestServiceBuilder struct {

}

func (t TestServiceBuilder) BuildArticleFactory(s *procezor.ProcezorService) bool {
	s.ArticleFactory = NewTestArticleFactory()
	if s.ArticleFactory == nil {
		return false
	}
	s.ArticleFactory.BuildFactory()
	return true
}

func (t TestServiceBuilder) BuildConceptFactory(s *procezor.ProcezorService) bool {
	s.ConceptFactory = NewTestConceptFactory()
	if s.ConceptFactory == nil {
		return false
	}
	s.ConceptFactory.BuildFactory()
	return true
}

type TestService struct {
	procezor.IProcezorService
}

func NewTestServiceBuilder() procezor.IProcezorFactoryBuilder {
	return &TestServiceBuilder{}
}

func NewTestService() procezor.IProcezorService{
	const (
		TestVersion      = TEST_VERSION
		TestFinalArticle = ARTICLE_INCOME_NETTO
		TestFinalConcept = CONCEPT_INCOME_NETTO
	)
	var (
		TestFinalDefs = procezor.GetArticleDefine(TestFinalArticle.Id(), TestFinalConcept.Id())
	)

	return &TestService{
		procezor.NewProcezorService(TestVersion, TestFinalDefs, NewTestServiceBuilder()),
	}
}

func GetTargetsWithSalaryHomeOffice(period legalios.IPeriod) types.ITermTargetList {
	const CONTRACT_CODE = 0
	const POSITION_CODE = 0
	const VARIANT1_CODE = 1

	var montCode = types.GetMonthCode(period.GetCode())
	var contract = types.GetContractCode(CONTRACT_CODE)
	var position = types.GetPositionCode(POSITION_CODE)
	var variant1 = types.GetVariantCode(VARIANT1_CODE)

	var targets = []types.ITermTarget {
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING.Id()), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY.Id()),procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_ALLOWCE_HOFFICE.Id()),procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
	}
	return targets
}

func GetTargetsWithSalaryBonusBarter(period legalios.IPeriod) types.ITermTargetList {
	const CONTRACT_CODE = 0
	const POSITION_CODE = 0
	const VARIANT1_CODE = 1

	var montCode = types.GetMonthCode(period.GetCode())
	var contract = types.GetContractCode(CONTRACT_CODE)
	var position = types.GetPositionCode(POSITION_CODE)
	var variant1 = types.GetVariantCode(VARIANT1_CODE)

	var targets = []types.ITermTarget {
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_TIMESHT_WORKING.Id()), procezor.GetConceptCode(CONCEPT_TIMESHT_WORKING.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_SALARY.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_BASIS.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BONUS.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
		types.NewTermTarget(montCode, contract, position, variant1,
			procezor.GetArticleCode(ARTICLE_PAYMENT_BARTER.Id()), procezor.GetConceptCode(CONCEPT_AMOUNT_FIXED.Id())),
	}
	return targets
}

