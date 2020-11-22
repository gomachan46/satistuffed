package data

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
)

type Data struct {
	Items *[]model.Item
}

func (d Data) GetItem(name string) (*model.Item, error) {
	for _, v := range *d.Items {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Error: item not found %s", name)
}

func (d Data) GetItemNames() []string {
	itemNames := []string{}

	for _, v := range *d.Items {
		itemNames = append(itemNames, v.Name)
	}

	return itemNames
}

func Load() *Data {
	ironOre := &model.Item{Name: "鉄鉱石"}
	copperOre := &model.Item{Name: "銅鉱石"}
	limestone := &model.Item{Name: "石灰岩"}
	coal := &model.Item{Name: "石炭"}
	cateriumOre := &model.Item{Name: "カテリウム鉱石"}
	rawQuartz := &model.Item{Name: "未加工石英"}
	sulfur := &model.Item{Name: "硫黄"}
	bauxite := &model.Item{Name: "ボーキサイト"}
	uranium := &model.Item{Name: "ウラン"}
	samOre := &model.Item{Name: "S.A.M.鉱石"}
	crudeOil := &model.Item{Name: "原油"}
	water := &model.Item{Name: "水"}

	leaves := &model.Item{Name: "葉"}
	wood := &model.Item{Name: "木"}
	flowerPetals := &model.Item{Name: "花弁"}
	mycelia := &model.Item{Name: "菌糸"}
	alienCarapace := &model.Item{Name: "異星生物甲殻"}
	alienOrgans := &model.Item{Name: "異星生物器官"}
	greenPowerSlug := &model.Item{Name: "緑のパワー・スラッグ"}
	yellowPowerSlug := &model.Item{Name: "黄のパワー・スラッグ"}
	purplePowerSlug := &model.Item{Name: "紫のパワー・スラッグ"}

	ironIngot := &model.Item{Name: "鉄のインゴット"}
	copperIngot := &model.Item{Name: "銅のインゴット"}
	steelIngot := &model.Item{Name: "鋼鉄のインゴット"}
	cateriumIngot := &model.Item{Name: "カテリウムのインゴット"}
	aluminumIngot := &model.Item{Name: "アルミのインゴット"}

	ironPlate := &model.Item{Name: "鉄板"}
	ironRod := &model.Item{Name: "鉄のロッド"}
	screw := &model.Item{Name: "ネジ"}
	reinforcedIronPlate := &model.Item{Name: "強化鉄板"}
	modularFrame := &model.Item{Name: "モジュラー・フレーム"}
	steelBeam := &model.Item{Name: "鋼梁"}
	steelPipe := &model.Item{Name: "鋼鉄のパイプ"}
	encasedIndustrialBeam := &model.Item{Name: "コンクリート被覆型鋼梁"}
	heavyModularFrame := &model.Item{Name: "ヘビー・モジュラー・フレーム"}
	alcladAluminumSheet := &model.Item{Name: "アルクラッド・アルミ板"}
	copperSheet := &model.Item{Name: "銅のシート"}

	wire := &model.Item{Name: "ワイヤー"}
	cable := &model.Item{Name: "ケーブル"}
	quickwire := &model.Item{Name: "クイックワイヤー"}
	circuitBoard := &model.Item{Name: "回路基板"}
	aiLimiter := &model.Item{Name: "A.I.リミッター"}
	highSpeedConnector := &model.Item{Name: "高速コネクター"}
	battery := &model.Item{Name: "バッテリー"}

	biomass := &model.Item{Name: "バイオマス"}
	solidBiofuel := &model.Item{Name: "固体バイオ燃料"}
	fabric := &model.Item{Name: "布地"}

	concrete := &model.Item{Name: "コンクリート"}
	quartzCrystal := &model.Item{Name: "石英結晶"}
	silica := &model.Item{Name: "シリカ"}
	compactedCoal := &model.Item{Name: "圧縮石炭"}

	rotor := &model.Item{Name: "ローター"}
	stator := &model.Item{Name: "固定子"}
	motor := &model.Item{Name: "モーター"}
	heatSink := &model.Item{Name: "ヒートシンク"}
	turboMotor := &model.Item{Name: "ターボ・モーター"}

	powerShard := &model.Item{Name: "パワー・シャード"}

	plastic := &model.Item{Name: "プラスチック"}
	rubber := &model.Item{Name: "ゴム"}
	petroleumCoke := &model.Item{Name: "石油コークス"}
	polymerResin := &model.Item{Name: "合成樹脂"}

	fuel := &model.Item{Name: "燃料"}
	turbofuel := &model.Item{Name: "ターボ燃料"}
	liquidBiofuel := &model.Item{Name: "液体バイオ燃料"}

	packagedFuel := &model.Item{Name: "容器入り燃料"}
	packagedTurbofuel := &model.Item{Name: "容器入りターボ燃料"}
	packagedLiquidBiofuel := &model.Item{Name: "容器入り液体バイオ燃料"}
	packagedHeavyOilResidue := &model.Item{Name: "容器入り廃重油"}
	packagedOil := &model.Item{Name: "容器入り原油"}
	packagedWater := &model.Item{Name: "容器入り水"}

	aluminumScrap := &model.Item{Name: "アルミのスクラップ"}
	aluminaSolution := &model.Item{Name: "アルミナ溶液"}
	sulfuricAcid := &model.Item{Name: "硫酸"}
	uraniumPellet := &model.Item{Name: "ウランペレット"}

	emptyCanister := &model.Item{Name: "空の容器"}
	heavyOilResidue := &model.Item{Name: "廃重油"}

	smartPlating := &model.Item{Name: "スマート・プレート"}
	versatileFramework := &model.Item{Name: "多目的フレームワーク"}
	automatedWiring := &model.Item{Name: "自動ワイヤー"}
	modularEngine := &model.Item{Name: "モジュラーエンジン"}
	adaptiveControlUnit := &model.Item{Name: "自律制御ユニット"}

	crystalOscillator := &model.Item{Name: "水晶発振器"}
	computer := &model.Item{Name: "コンピューター"}
	supercomputer := &model.Item{Name: "スーパーコンピューター"}
	radioControlUnit := &model.Item{Name: "無線制御ユニット"}

	encasedUraniumCell := &model.Item{Name: "被覆型ウラン・セル"}
	electromagneticControlRod := &model.Item{Name: "電磁制御棒"}
	nuclearFuelRod := &model.Item{Name: "核燃料棒"}

	ironOreRecipe := &model.Recipe{Name: "鉄鉱石のレシピ1"}
	copperOreRecipe := &model.Recipe{Name: "銅鉱石のレシピ1"}
	limestoneRecipe := &model.Recipe{Name: "石灰岩のレシピ1"}
	coalRecipe := &model.Recipe{Name: "石炭のレシピ1"}
	cateriumOreRecipe := &model.Recipe{Name: "カテリウム鉱石のレシピ1"}
	rawQuartzRecipe := &model.Recipe{Name: "未加工石英のレシピ1"}
	sulfurRecipe := &model.Recipe{Name: "硫黄のレシピ1"}
	bauxiteRecipe := &model.Recipe{Name: "ボーキサイトのレシピ1"}
	uraniumRecipe := &model.Recipe{Name: "ウランのレシピ1"}
	samOreRecipe := &model.Recipe{Name: "S.A.M.鉱石のレシピ1"}
	crudeOilRecipe := &model.Recipe{Name: "原油のレシピ1"}
	waterRecipe := &model.Recipe{Name: "水のレシピ1"}

	leavesRecipe := &model.Recipe{Name: "葉のレシピ1"}
	woodRecipe := &model.Recipe{Name: "木のレシピ1"}
	flowerPetalsRecipe := &model.Recipe{Name: "花弁のレシピ1"}
	myceliaRecipe := &model.Recipe{Name: "菌糸のレシピ1"}
	alienCarapaceRecipe := &model.Recipe{Name: "異星生物甲殻のレシピ1"}
	alienOrgansRecipe := &model.Recipe{Name: "異星生物器官のレシピ1"}
	greenPowerSlugRecipe := &model.Recipe{Name: "緑のパワー・スラッグのレシピ1"}
	yellowPowerSlugRecipe := &model.Recipe{Name: "黄のパワー・スラッグのレシピ1"}
	purplePowerSlugRecipe := &model.Recipe{Name: "紫のパワー・スラッグのレシピ1"}

	ironIngotRecipe := &model.Recipe{Name: "鉄インゴットのレシピ1"}
	copperIngotRecipe := &model.Recipe{Name: "銅のインゴットのレシピ1"}
	steelIngotRecipe := &model.Recipe{Name: "鋼鉄のインゴットのレシピ1"}
	cateriumIngotRecipe := &model.Recipe{Name: "カテリウムのインゴットのレシピ1"}
	aluminumIngotRecipe := &model.Recipe{Name: "アルミのインゴットのレシピ1"}

	ironPlateRecipe := &model.Recipe{Name: "鉄板のレシピ1"}
	ironRodRecipe := &model.Recipe{Name: "鉄のロッドのレシピ1"}
	screwRecipe := &model.Recipe{Name: "ネジのレシピ1"}
	reinforcedIronPlateRecipe := &model.Recipe{Name: "強化鉄板のレシピ1"}
	modularFrameRecipe := &model.Recipe{Name: "モジュラー・フレームのレシピ1"}
	steelBeamRecipe := &model.Recipe{Name: "鋼梁のレシピ1"}
	steelPipeRecipe := &model.Recipe{Name: "鋼鉄のパイプのレシピ1"}
	encasedIndustrialBeamRecipe := &model.Recipe{Name: "コンクリート被覆型鋼梁のレシピ1"}
	heavyModularFrameRecipe := &model.Recipe{Name: "ヘビー・モジュラー・フレームのレシピ1"}
	alcladAluminumSheetRecipe := &model.Recipe{Name: "アルクラッド・アルミ板のレシピ1"}
	copperSheetRecipe := &model.Recipe{Name: "銅のシートのレシピ1"}

	wireRecipe := &model.Recipe{Name: "ワイヤーのレシピ1"}
	cableRecipe := &model.Recipe{Name: "ケーブルのレシピ1"}
	quickwireRecipe := &model.Recipe{Name: "クイックワイヤーのレシピ1"}
	circuitBoardRecipe := &model.Recipe{Name: "回路基板のレシピ1"}
	aiLimiterRecipe := &model.Recipe{Name: "A.I.リミッターのレシピ1"}
	highSpeedConnectorRecipe := &model.Recipe{Name: "高速コネクターのレシピ1"}
	batteryRecipe := &model.Recipe{Name: "バッテリーのレシピ1"}

	biomassRecipe := &model.Recipe{Name: "バイオマスのレシピ1"}
	solidBiofuelRecipe := &model.Recipe{Name: "固体バイオ燃料のレシピ1"}
	fabricRecipe := &model.Recipe{Name: "布地のレシピ1"}

	concreteRecipe := &model.Recipe{Name: "コンクリートのレシピ1"}
	quartzCrystalRecipe := &model.Recipe{Name: "石英結晶のレシピ1"}
	silicaRecipe := &model.Recipe{Name: "シリカのレシピ1"}
	compactedCoalRecipe := &model.Recipe{Name: "圧縮石炭のレシピ1"}

	rotorRecipe := &model.Recipe{Name: "ローターのレシピ1"}
	statorRecipe := &model.Recipe{Name: "固定子のレシピ1"}
	motorRecipe := &model.Recipe{Name: "モーターのレシピ1"}
	heatSinkRecipe := &model.Recipe{Name: "ヒートシンクのレシピ1"}
	turboMotorRecipe := &model.Recipe{Name: "ターボ・モーターのレシピ1"}

	powerShardRecipe := &model.Recipe{Name: "パワー・シャードのレシピ1"}

	plasticRecipe := &model.Recipe{Name: "プラスチックのレシピ1"}
	rubberRecipe := &model.Recipe{Name: "ゴムのレシピ1"}
	petroleumCokeRecipe := &model.Recipe{Name: "石油コークスのレシピ1"}
	//polymerResinRecipe := &model.Recipe{Name: "合成樹脂のレシピ1"}

	fuelRecipe := &model.Recipe{Name: "燃料のレシピ1"}
	turbofuelRecipe := &model.Recipe{Name: "ターボ燃料のレシピ1"}
	liquidBiofuelRecipe := &model.Recipe{Name: "液体バイオ燃料のレシピ1"}

	packagedFuelRecipe := &model.Recipe{Name: "容器入り燃料のレシピ1"}
	packagedTurbofuelRecipe := &model.Recipe{Name: "容器入りターボ燃料のレシピ1"}
	packagedLiquidBiofuelRecipe := &model.Recipe{Name: "容器入り液体バイオ燃料のレシピ1"}
	packagedHeavyOilResidueRecipe := &model.Recipe{Name: "容器入り廃重油のレシピ1"}
	packagedOilRecipe := &model.Recipe{Name: "容器入り原油のレシピ1"}
	packagedWaterRecipe := &model.Recipe{Name: "容器入り水のレシピ1"}

	aluminumScrapRecipe := &model.Recipe{Name: "アルミのスクラップのレシピ1"}
	aluminaSolutionRecipe := &model.Recipe{Name: "アルミナ溶液のレシピ1"}
	sulfuricAcidRecipe := &model.Recipe{Name: "硫酸のレシピ1"}
	uraniumPelletRecipe := &model.Recipe{Name: "ウランペレットのレシピ1"}

	emptyCanisterRecipe := &model.Recipe{Name: "空の容器のレシピ1"}
	heavyOilResidueRecipe := &model.Recipe{Name: "廃重油のレシピ1"}

	smartPlatingRecipe := &model.Recipe{Name: "スマート・プレートのレシピ1"}
	versatileFrameworkRecipe := &model.Recipe{Name: "多目的フレームワークのレシピ1"}
	automatedWiringRecipe := &model.Recipe{Name: "自動ワイヤーのレシピ1"}
	modularEngineRecipe := &model.Recipe{Name: "モジュラーエンジンのレシピ1"}
	adaptiveControlUnitRecipe := &model.Recipe{Name: "自律制御ユニットのレシピ1"}

	crystalOscillatorRecipe := &model.Recipe{Name: "水晶発振器のレシピ1"}
	computerRecipe := &model.Recipe{Name: "コンピューターのレシピ1"}
	supercomputerRecipe := &model.Recipe{Name: "スーパーコンピューターのレシピ1"}
	radioControlUnitRecipe := &model.Recipe{Name: "無線制御ユニットのレシピ1"}

	encasedUraniumCellRecipe := &model.Recipe{Name: "被覆型ウラン・セルのレシピ1"}
	electromagneticControlRodRecipe := &model.Recipe{Name: "電磁制御棒のレシピ1"}
	nuclearFuelRodRecipe := &model.Recipe{Name: "核燃料棒のレシピ1"}

	ironOreRecipe.Ingredients = &[]model.Ingredient{}
	ironOreRecipe.Products = &[]model.Product{{Item: ironOre, Amount: 60}}
	ironOre.Recipes = &[]model.Recipe{*ironOreRecipe}
	copperOreRecipe.Ingredients = &[]model.Ingredient{}
	copperOreRecipe.Products = &[]model.Product{{Item: copperOre, Amount: 60}}
	copperOre.Recipes = &[]model.Recipe{*copperOreRecipe}
	limestoneRecipe.Ingredients = &[]model.Ingredient{}
	limestoneRecipe.Products = &[]model.Product{{Item: limestone, Amount: 60}}
	limestone.Recipes = &[]model.Recipe{*limestoneRecipe}
	coalRecipe.Ingredients = &[]model.Ingredient{}
	coalRecipe.Products = &[]model.Product{{Item: coal, Amount: 60}}
	coal.Recipes = &[]model.Recipe{*coalRecipe}
	cateriumOreRecipe.Ingredients = &[]model.Ingredient{}
	cateriumOreRecipe.Products = &[]model.Product{{Item: cateriumOre, Amount: 60}}
	cateriumOre.Recipes = &[]model.Recipe{*cateriumOreRecipe}
	rawQuartzRecipe.Ingredients = &[]model.Ingredient{}
	rawQuartzRecipe.Products = &[]model.Product{{Item: rawQuartz, Amount: 60}}
	rawQuartz.Recipes = &[]model.Recipe{*rawQuartzRecipe}
	sulfurRecipe.Ingredients = &[]model.Ingredient{}
	sulfurRecipe.Products = &[]model.Product{{Item: sulfur, Amount: 60}}
	sulfur.Recipes = &[]model.Recipe{*sulfurRecipe}
	bauxiteRecipe.Ingredients = &[]model.Ingredient{}
	bauxiteRecipe.Products = &[]model.Product{{Item: bauxite, Amount: 60}}
	bauxite.Recipes = &[]model.Recipe{*bauxiteRecipe}
	uraniumRecipe.Ingredients = &[]model.Ingredient{}
	uraniumRecipe.Products = &[]model.Product{{Item: uranium, Amount: 60}}
	uranium.Recipes = &[]model.Recipe{*uraniumRecipe}
	samOreRecipe.Ingredients = &[]model.Ingredient{}
	samOreRecipe.Products = &[]model.Product{{Item: samOre, Amount: 60}}
	samOre.Recipes = &[]model.Recipe{*samOreRecipe}
	crudeOilRecipe.Ingredients = &[]model.Ingredient{}
	crudeOilRecipe.Products = &[]model.Product{{Item: crudeOil, Amount: 60}}
	crudeOil.Recipes = &[]model.Recipe{*crudeOilRecipe}
	waterRecipe.Ingredients = &[]model.Ingredient{}
	waterRecipe.Products = &[]model.Product{{Item: water, Amount: 60}}
	water.Recipes = &[]model.Recipe{*waterRecipe}

	leavesRecipe.Ingredients = &[]model.Ingredient{}
	leavesRecipe.Products = &[]model.Product{{Item: leaves, Amount: 60}}
	leaves.Recipes = &[]model.Recipe{*leavesRecipe}
	woodRecipe.Ingredients = &[]model.Ingredient{}
	woodRecipe.Products = &[]model.Product{{Item: wood, Amount: 60}}
	wood.Recipes = &[]model.Recipe{*woodRecipe}
	flowerPetalsRecipe.Ingredients = &[]model.Ingredient{}
	flowerPetalsRecipe.Products = &[]model.Product{{Item: flowerPetals, Amount: 60}}
	flowerPetals.Recipes = &[]model.Recipe{*flowerPetalsRecipe}
	myceliaRecipe.Ingredients = &[]model.Ingredient{}
	myceliaRecipe.Products = &[]model.Product{{Item: mycelia, Amount: 60}}
	mycelia.Recipes = &[]model.Recipe{*myceliaRecipe}
	alienCarapaceRecipe.Ingredients = &[]model.Ingredient{}
	alienCarapaceRecipe.Products = &[]model.Product{{Item: alienCarapace, Amount: 60}}
	alienCarapace.Recipes = &[]model.Recipe{*alienCarapaceRecipe}
	alienOrgansRecipe.Ingredients = &[]model.Ingredient{}
	alienOrgansRecipe.Products = &[]model.Product{{Item: alienOrgans, Amount: 60}}
	alienOrgans.Recipes = &[]model.Recipe{*alienOrgansRecipe}
	greenPowerSlugRecipe.Ingredients = &[]model.Ingredient{}
	greenPowerSlugRecipe.Products = &[]model.Product{{Item: greenPowerSlug, Amount: 60}}
	greenPowerSlug.Recipes = &[]model.Recipe{*greenPowerSlugRecipe}
	yellowPowerSlugRecipe.Ingredients = &[]model.Ingredient{}
	yellowPowerSlugRecipe.Products = &[]model.Product{{Item: yellowPowerSlug, Amount: 60}}
	yellowPowerSlug.Recipes = &[]model.Recipe{*yellowPowerSlugRecipe}
	purplePowerSlugRecipe.Ingredients = &[]model.Ingredient{}
	purplePowerSlugRecipe.Products = &[]model.Product{{Item: purplePowerSlug, Amount: 60}}
	purplePowerSlug.Recipes = &[]model.Recipe{*purplePowerSlugRecipe}

	ironIngotRecipe.Ingredients = &[]model.Ingredient{{Item: ironOre, Amount: 30}}
	ironIngotRecipe.Products = &[]model.Product{{Item: ironIngot, Amount: 30}}
	ironIngot.Recipes = &[]model.Recipe{*ironIngotRecipe}
	copperIngotRecipe.Ingredients = &[]model.Ingredient{{Item: ironOre, Amount: 30}}
	copperIngotRecipe.Products = &[]model.Product{{Item: copperIngot, Amount: 30}}
	copperIngot.Recipes = &[]model.Recipe{*copperIngotRecipe}
	steelIngotRecipe.Ingredients = &[]model.Ingredient{{Item: ironOre, Amount: 45}, {Item: coal, Amount: 45}}
	steelIngotRecipe.Products = &[]model.Product{{Item: steelIngot, Amount: 45}}
	steelIngot.Recipes = &[]model.Recipe{*steelIngotRecipe}
	cateriumIngotRecipe.Ingredients = &[]model.Ingredient{{Item: cateriumOre, Amount: 45}}
	cateriumIngotRecipe.Products = &[]model.Product{{Item: cateriumIngot, Amount: 15}}
	cateriumIngot.Recipes = &[]model.Recipe{*cateriumIngotRecipe}
	aluminumIngotRecipe.Ingredients = &[]model.Ingredient{{Item: aluminumScrap, Amount: 240}, {Item: silica, Amount: 140}}
	aluminumIngotRecipe.Products = &[]model.Product{{Item: aluminumIngot, Amount: 80}}
	aluminumIngot.Recipes = &[]model.Recipe{*aluminumIngotRecipe}

	concreteRecipe.Ingredients = &[]model.Ingredient{{Item: limestone, Amount: 45}}
	concreteRecipe.Products = &[]model.Product{{Item: concrete, Amount: 15}}
	concrete.Recipes = &[]model.Recipe{*concreteRecipe}
	quartzCrystalRecipe.Ingredients = &[]model.Ingredient{{Item: rawQuartz, Amount: 37.5}}
	quartzCrystalRecipe.Products = &[]model.Product{{Item: quartzCrystal, Amount: 22.5}}
	quartzCrystal.Recipes = &[]model.Recipe{*quartzCrystalRecipe}
	silicaRecipe.Ingredients = &[]model.Ingredient{{Item: rawQuartz, Amount: 22.5}}
	silicaRecipe.Products = &[]model.Product{{Item: silica, Amount: 37.5}}
	silica.Recipes = &[]model.Recipe{*silicaRecipe}
	compactedCoalRecipe.Ingredients = &[]model.Ingredient{{Item: coal, Amount: 25}, {Item: sulfur, Amount: 25}}
	compactedCoalRecipe.Products = &[]model.Product{{Item: compactedCoal, Amount: 25}}
	compactedCoal.Recipes = &[]model.Recipe{*compactedCoalRecipe}

	plasticRecipe.Ingredients = &[]model.Ingredient{{Item: crudeOil, Amount: 30}}
	plasticRecipe.Products = &[]model.Product{{Item: plastic, Amount: 20}, {Item: heavyOilResidue, Amount: 10}}
	plastic.Recipes = &[]model.Recipe{*plasticRecipe}
	rubberRecipe.Ingredients = &[]model.Ingredient{{Item: crudeOil, Amount: 30}}
	rubberRecipe.Products = &[]model.Product{{Item: rubber, Amount: 20}, {Item: heavyOilResidue, Amount: 20}}
	rubber.Recipes = &[]model.Recipe{*rubberRecipe}
	petroleumCokeRecipe.Ingredients = &[]model.Ingredient{{Item: heavyOilResidue, Amount: 40}}
	petroleumCokeRecipe.Products = &[]model.Product{{Item: petroleumCoke, Amount: 120}}
	petroleumCoke.Recipes = &[]model.Recipe{*petroleumCokeRecipe}

	fuelRecipe.Ingredients = &[]model.Ingredient{{Item: crudeOil, Amount: 60}}
	fuelRecipe.Products = &[]model.Product{{Item: fuel, Amount: 40}, {Item: polymerResin, Amount: 30}}
	fuel.Recipes = &[]model.Recipe{*fuelRecipe}
	turbofuelRecipe.Ingredients = &[]model.Ingredient{{Item: fuel, Amount: 22.5}, {Item: compactedCoal, Amount: 15}}
	turbofuelRecipe.Products = &[]model.Product{{Item: turbofuel, Amount: 18.75}}
	turbofuel.Recipes = &[]model.Recipe{*turbofuelRecipe}
	liquidBiofuelRecipe.Ingredients = &[]model.Ingredient{{Item: solidBiofuel, Amount: 90}, {Item: water, Amount: 45}}
	liquidBiofuelRecipe.Products = &[]model.Product{{Item: liquidBiofuel, Amount: 60}}
	liquidBiofuel.Recipes = &[]model.Recipe{*liquidBiofuelRecipe}

	packagedFuelRecipe.Ingredients = &[]model.Ingredient{{Item: fuel, Amount: 40}, {Item: emptyCanister, Amount: 40}}
	packagedFuelRecipe.Products = &[]model.Product{{Item: packagedFuel, Amount: 40}}
	packagedFuel.Recipes = &[]model.Recipe{*packagedFuelRecipe}
	packagedTurbofuelRecipe.Ingredients = &[]model.Ingredient{{Item: turbofuel, Amount: 20}, {Item: emptyCanister, Amount: 20}}
	packagedTurbofuelRecipe.Products = &[]model.Product{{Item: packagedTurbofuel, Amount: 20}}
	packagedTurbofuel.Recipes = &[]model.Recipe{*packagedTurbofuelRecipe}
	packagedLiquidBiofuelRecipe.Ingredients = &[]model.Ingredient{{Item: liquidBiofuel, Amount: 40}, {Item: emptyCanister, Amount: 40}}
	packagedLiquidBiofuelRecipe.Products = &[]model.Product{{Item: packagedLiquidBiofuel, Amount: 40}}
	packagedLiquidBiofuel.Recipes = &[]model.Recipe{*packagedLiquidBiofuelRecipe}
	packagedHeavyOilResidueRecipe.Ingredients = &[]model.Ingredient{{Item: heavyOilResidue, Amount: 30}, {Item: emptyCanister, Amount: 30}}
	packagedHeavyOilResidueRecipe.Products = &[]model.Product{{Item: packagedHeavyOilResidue, Amount: 30}}
	packagedHeavyOilResidue.Recipes = &[]model.Recipe{*packagedHeavyOilResidueRecipe}
	packagedOilRecipe.Ingredients = &[]model.Ingredient{{Item: crudeOil, Amount: 30}, {Item: emptyCanister, Amount: 30}}
	packagedOilRecipe.Products = &[]model.Product{{Item: packagedOil, Amount: 30}}
	packagedOil.Recipes = &[]model.Recipe{*packagedOilRecipe}
	packagedWaterRecipe.Ingredients = &[]model.Ingredient{{Item: water, Amount: 60}, {Item: emptyCanister, Amount: 60}}
	packagedWaterRecipe.Products = &[]model.Product{{Item: packagedWater, Amount: 60}}
	packagedWater.Recipes = &[]model.Recipe{*packagedWaterRecipe}

	aluminumScrapRecipe.Ingredients = &[]model.Ingredient{{Item: aluminaSolution, Amount: 240}, {Item: petroleumCoke, Amount: 60}}
	aluminumScrapRecipe.Products = &[]model.Product{{Item: aluminumScrap, Amount: 360}, {Item: water, Amount: 60}}
	aluminumScrap.Recipes = &[]model.Recipe{*aluminumScrapRecipe}
	aluminaSolutionRecipe.Ingredients = &[]model.Ingredient{{Item: bauxite, Amount: 70}, {Item: water, Amount: 100}}
	aluminaSolutionRecipe.Products = &[]model.Product{{Item: aluminaSolution, Amount: 80}, {Item: silica, Amount: 20}}
	aluminaSolution.Recipes = &[]model.Recipe{*aluminaSolutionRecipe}
	sulfuricAcidRecipe.Ingredients = &[]model.Ingredient{{Item: sulfur, Amount: 50}, {Item: water, Amount: 50}}
	sulfuricAcidRecipe.Products = &[]model.Product{{Item: sulfuricAcid, Amount: 100}}
	sulfuricAcid.Recipes = &[]model.Recipe{*sulfuricAcidRecipe}
	uraniumPelletRecipe.Ingredients = &[]model.Ingredient{{Item: uranium, Amount: 50}, {Item: sulfuricAcid, Amount: 80}}
	uraniumPelletRecipe.Products = &[]model.Product{{Item: uraniumPellet, Amount: 50}, {Item: sulfuricAcid, Amount: 20}}
	uraniumPellet.Recipes = &[]model.Recipe{*uraniumPelletRecipe}

	emptyCanisterRecipe.Ingredients = &[]model.Ingredient{{Item: plastic, Amount: 30}}
	emptyCanisterRecipe.Products = &[]model.Product{{Item: emptyCanister, Amount: 60}}
	emptyCanister.Recipes = &[]model.Recipe{*emptyCanisterRecipe}
	// TODO: ゴムの副産物も対応したい
	heavyOilResidueRecipe.Ingredients = &[]model.Ingredient{{Item: crudeOil, Amount: 30}}
	heavyOilResidueRecipe.Products = &[]model.Product{{Item: heavyOilResidue, Amount: 10}, {Item: plastic, Amount: 20}}
	heavyOilResidue.Recipes = &[]model.Recipe{*heavyOilResidueRecipe}

	biomassRecipe.Ingredients = &[]model.Ingredient{{Item: leaves, Amount: 120}}
	biomassRecipe.Products = &[]model.Product{{Item: biomass, Amount: 60}}
	biomass.Recipes = &[]model.Recipe{*biomassRecipe}
	solidBiofuelRecipe.Ingredients = &[]model.Ingredient{{Item: biomass, Amount: 120}}
	solidBiofuelRecipe.Products = &[]model.Product{{Item: solidBiofuel, Amount: 60}}
	solidBiofuel.Recipes = &[]model.Recipe{*solidBiofuelRecipe}
	fabricRecipe.Ingredients = &[]model.Ingredient{{Item: mycelia, Amount: 15}, {Item: biomass, Amount: 75}}
	fabricRecipe.Products = &[]model.Product{{Item: fabric, Amount: 15}}
	fabric.Recipes = &[]model.Recipe{*fabricRecipe}

	ironPlateRecipe.Ingredients = &[]model.Ingredient{{Item: ironIngot, Amount: 30}}
	ironPlateRecipe.Products = &[]model.Product{{Item: ironPlate, Amount: 20}}
	ironPlate.Recipes = &[]model.Recipe{*ironPlateRecipe}
	ironRodRecipe.Ingredients = &[]model.Ingredient{{Item: ironIngot, Amount: 15}}
	ironRodRecipe.Products = &[]model.Product{{Item: ironRod, Amount: 15}}
	ironRod.Recipes = &[]model.Recipe{*ironRodRecipe}
	screwRecipe.Ingredients = &[]model.Ingredient{{Item: ironRod, Amount: 10}}
	screwRecipe.Products = &[]model.Product{{Item: screw, Amount: 40}}
	screw.Recipes = &[]model.Recipe{*screwRecipe}
	reinforcedIronPlateRecipe.Ingredients = &[]model.Ingredient{{Item: ironPlate, Amount: 30}, {Item: screw, Amount: 60}}
	reinforcedIronPlateRecipe.Products = &[]model.Product{{Item: reinforcedIronPlate, Amount: 5}}
	reinforcedIronPlate.Recipes = &[]model.Recipe{*reinforcedIronPlateRecipe}
	modularFrameRecipe.Ingredients = &[]model.Ingredient{{Item: reinforcedIronPlate, Amount: 3}, {Item: ironRod, Amount: 12}}
	modularFrameRecipe.Products = &[]model.Product{{Item: modularFrame, Amount: 2}}
	modularFrame.Recipes = &[]model.Recipe{*modularFrameRecipe}
	steelBeamRecipe.Ingredients = &[]model.Ingredient{{Item: steelIngot, Amount: 60}}
	steelBeamRecipe.Products = &[]model.Product{{Item: steelBeam, Amount: 15}}
	steelBeam.Recipes = &[]model.Recipe{*steelBeamRecipe}
	steelPipeRecipe.Ingredients = &[]model.Ingredient{{Item: steelIngot, Amount: 30}}
	steelPipeRecipe.Products = &[]model.Product{{Item: steelPipe, Amount: 20}}
	steelPipe.Recipes = &[]model.Recipe{*steelPipeRecipe}
	encasedIndustrialBeamRecipe.Ingredients = &[]model.Ingredient{{Item: steelBeam, Amount: 24}, {Item: concrete, Amount: 30}}
	encasedIndustrialBeamRecipe.Products = &[]model.Product{{Item: encasedIndustrialBeam, Amount: 6}}
	encasedIndustrialBeam.Recipes = &[]model.Recipe{*encasedIndustrialBeamRecipe}
	heavyModularFrameRecipe.Ingredients = &[]model.Ingredient{{Item: modularFrame, Amount: 10}, {Item: steelPipe, Amount: 30}, {Item: encasedIndustrialBeam, Amount: 10}, {Item: screw, Amount: 200}}
	heavyModularFrameRecipe.Products = &[]model.Product{{Item: heavyModularFrame, Amount: 2}}
	heavyModularFrame.Recipes = &[]model.Recipe{*heavyModularFrameRecipe}
	alcladAluminumSheetRecipe.Ingredients = &[]model.Ingredient{{Item: aluminumIngot, Amount: 60}, {Item: copperIngot, Amount: 22.5}}
	alcladAluminumSheetRecipe.Products = &[]model.Product{{Item: alcladAluminumSheet, Amount: 30}}
	alcladAluminumSheet.Recipes = &[]model.Recipe{*alcladAluminumSheetRecipe}
	copperSheetRecipe.Ingredients = &[]model.Ingredient{{Item: copperIngot, Amount: 20}}
	copperSheetRecipe.Products = &[]model.Product{{Item: copperSheet, Amount: 10}}
	copperSheet.Recipes = &[]model.Recipe{*copperSheetRecipe}

	rotorRecipe.Ingredients = &[]model.Ingredient{{Item: ironRod, Amount: 20}, {Item: screw, Amount: 100}}
	rotorRecipe.Products = &[]model.Product{{Item: rotor, Amount: 4}}
	rotor.Recipes = &[]model.Recipe{*rotorRecipe}
	statorRecipe.Ingredients = &[]model.Ingredient{{Item: steelPipe, Amount: 15}, {Item: wire, Amount: 40}}
	statorRecipe.Products = &[]model.Product{{Item: stator, Amount: 5}}
	stator.Recipes = &[]model.Recipe{*statorRecipe}
	motorRecipe.Ingredients = &[]model.Ingredient{{Item: rotor, Amount: 10}, {Item: stator, Amount: 10}}
	motorRecipe.Products = &[]model.Product{{Item: motor, Amount: 5}}
	motor.Recipes = &[]model.Recipe{*motorRecipe}
	heatSinkRecipe.Ingredients = &[]model.Ingredient{{Item: alcladAluminumSheet, Amount: 40}, {Item: rubber, Amount: 70}}
	heatSinkRecipe.Products = &[]model.Product{{Item: heatSink, Amount: 10}}
	heatSink.Recipes = &[]model.Recipe{*heatSinkRecipe}
	turboMotorRecipe.Ingredients = &[]model.Ingredient{{Item: heatSink, Amount: 7.5}, {Item: radioControlUnit, Amount: 3.75}, {Item: motor, Amount: 7.5}, {Item: rubber, Amount: 45}}
	turboMotorRecipe.Products = &[]model.Product{{Item: turboMotor, Amount: 1.875}}
	turboMotor.Recipes = &[]model.Recipe{*turboMotorRecipe}

	powerShardRecipe.Ingredients = &[]model.Ingredient{{Item: greenPowerSlug, Amount: 1}}
	powerShardRecipe.Products = &[]model.Product{{Item: powerShard, Amount: 1}}
	powerShard.Recipes = &[]model.Recipe{*powerShardRecipe}

	wireRecipe.Ingredients = &[]model.Ingredient{{Item: copperIngot, Amount: 15}}
	wireRecipe.Products = &[]model.Product{{Item: wire, Amount: 30}}
	wire.Recipes = &[]model.Recipe{*wireRecipe}
	cableRecipe.Ingredients = &[]model.Ingredient{{Item: wire, Amount: 60}}
	cableRecipe.Products = &[]model.Product{{Item: cable, Amount: 30}}
	cable.Recipes = &[]model.Recipe{*cableRecipe}
	quickwireRecipe.Ingredients = &[]model.Ingredient{{Item: cateriumIngot, Amount: 12}}
	quickwireRecipe.Products = &[]model.Product{{Item: quickwire, Amount: 60}}
	quickwire.Recipes = &[]model.Recipe{*quickwireRecipe}
	circuitBoardRecipe.Ingredients = &[]model.Ingredient{{Item: copperSheet, Amount: 15}, {Item: plastic, Amount: 30}}
	circuitBoardRecipe.Products = &[]model.Product{{Item: circuitBoard, Amount: 7.5}}
	circuitBoard.Recipes = &[]model.Recipe{*circuitBoardRecipe}
	aiLimiterRecipe.Ingredients = &[]model.Ingredient{{Item: copperSheet, Amount: 25}, {Item: quickwire, Amount: 100}}
	aiLimiterRecipe.Products = &[]model.Product{{Item: aiLimiter, Amount: 5}}
	aiLimiter.Recipes = &[]model.Recipe{*aiLimiterRecipe}
	highSpeedConnectorRecipe.Ingredients = &[]model.Ingredient{{Item: quickwire, Amount: 210}, {Item: cable, Amount: 37.5}, {Item: circuitBoard, Amount: 3.75}}
	highSpeedConnectorRecipe.Products = &[]model.Product{{Item: highSpeedConnector, Amount: 3.75}}
	highSpeedConnector.Recipes = &[]model.Recipe{*highSpeedConnectorRecipe}
	batteryRecipe.Ingredients = &[]model.Ingredient{{Item: alcladAluminumSheet, Amount: 15}, {Item: wire, Amount: 30}, {Item: sulfur, Amount: 37.5}, {Item: plastic, Amount: 15}}
	batteryRecipe.Products = &[]model.Product{{Item: battery, Amount: 5.625}}
	battery.Recipes = &[]model.Recipe{*batteryRecipe}

	crystalOscillatorRecipe.Ingredients = &[]model.Ingredient{{Item: quartzCrystal, Amount: 18}, {Item: cable, Amount: 14}, {Item: reinforcedIronPlate, Amount: 2.5}}
	crystalOscillatorRecipe.Products = &[]model.Product{{Item: crystalOscillator, Amount: 1}}
	crystalOscillator.Recipes = &[]model.Recipe{*crystalOscillatorRecipe}
	computerRecipe.Ingredients = &[]model.Ingredient{{Item: circuitBoard, Amount: 25}, {Item: cable, Amount: 22.5}, {Item: plastic, Amount: 45}, {Item: screw, Amount: 130}}
	computerRecipe.Products = &[]model.Product{{Item: computer, Amount: 2.5}}
	computer.Recipes = &[]model.Recipe{*computerRecipe}
	supercomputerRecipe.Ingredients = &[]model.Ingredient{{Item: computer, Amount: 3.75}, {Item: aiLimiter, Amount: 3.75}, {Item: highSpeedConnector, Amount: 5.625}, {Item: plastic, Amount: 52.5}}
	supercomputerRecipe.Products = &[]model.Product{{Item: supercomputer, Amount: 1.875}}
	supercomputer.Recipes = &[]model.Recipe{*supercomputerRecipe}
	radioControlUnitRecipe.Ingredients = &[]model.Ingredient{{Item: heatSink, Amount: 10}, {Item: rubber, Amount: 40}, {Item: crystalOscillator, Amount: 2.5}, {Item: computer, Amount: 2.5}}
	radioControlUnitRecipe.Products = &[]model.Product{{Item: radioControlUnit, Amount: 2.5}}
	radioControlUnit.Recipes = &[]model.Recipe{*radioControlUnitRecipe}

	encasedUraniumCellRecipe.Ingredients = &[]model.Ingredient{{Item: uraniumPellet, Amount: 40}, {Item: concrete, Amount: 9}}
	encasedUraniumCellRecipe.Products = &[]model.Product{{Item: encasedUraniumCell, Amount: 10}}
	encasedUraniumCell.Recipes = &[]model.Recipe{*encasedUraniumCellRecipe}
	electromagneticControlRodRecipe.Ingredients = &[]model.Ingredient{{Item: stator, Amount: 6}, {Item: aiLimiter, Amount: 4}}
	electromagneticControlRodRecipe.Products = &[]model.Product{{Item: electromagneticControlRod, Amount: 4}}
	electromagneticControlRod.Recipes = &[]model.Recipe{*electromagneticControlRodRecipe}
	nuclearFuelRodRecipe.Ingredients = &[]model.Ingredient{{Item: encasedUraniumCell, Amount: 10}, {Item: encasedIndustrialBeam, Amount: 1.2}, {Item: electromagneticControlRod, Amount: 2}}
	nuclearFuelRodRecipe.Products = &[]model.Product{{Item: nuclearFuelRod, Amount: 0.4}}
	nuclearFuelRod.Recipes = &[]model.Recipe{*nuclearFuelRodRecipe}

	smartPlatingRecipe.Ingredients = &[]model.Ingredient{{Item: reinforcedIronPlate, Amount: 2}, {Item: rotor, Amount: 2}}
	smartPlatingRecipe.Products = &[]model.Product{{Item: smartPlating, Amount: 2}}
	smartPlating.Recipes = &[]model.Recipe{*smartPlatingRecipe}
	versatileFrameworkRecipe.Ingredients = &[]model.Ingredient{{Item: modularFrame, Amount: 2.5}, {Item: steelBeam, Amount: 30}}
	versatileFrameworkRecipe.Products = &[]model.Product{{Item: versatileFramework, Amount: 5}}
	versatileFramework.Recipes = &[]model.Recipe{*versatileFrameworkRecipe}
	automatedWiringRecipe.Ingredients = &[]model.Ingredient{{Item: stator, Amount: 2.5}, {Item: cable, Amount: 50}}
	automatedWiringRecipe.Products = &[]model.Product{{Item: automatedWiring, Amount: 2.5}}
	automatedWiring.Recipes = &[]model.Recipe{*automatedWiringRecipe}
	modularEngineRecipe.Ingredients = &[]model.Ingredient{{Item: motor, Amount: 2}, {Item: rubber, Amount: 15}, {Item: smartPlating, Amount: 2}}
	modularEngineRecipe.Products = &[]model.Product{{Item: modularEngine, Amount: 1}}
	modularEngine.Recipes = &[]model.Recipe{*modularEngineRecipe}
	adaptiveControlUnitRecipe.Ingredients = &[]model.Ingredient{{Item: automatedWiring, Amount: 7.5}, {Item: circuitBoard, Amount: 5}, {Item: heavyModularFrame, Amount: 1}, {Item: computer, Amount: 1}}
	adaptiveControlUnitRecipe.Products = &[]model.Product{{Item: adaptiveControlUnit, Amount: 1}}
	adaptiveControlUnit.Recipes = &[]model.Recipe{*adaptiveControlUnitRecipe}

	var items []model.Item
	items = append(
		items,
		*ironOre,
		*copperOre,
		*limestone,
		*coal,
		*cateriumOre,
		*rawQuartz,
		*sulfur,
		*bauxite,
		*uranium,
		*samOre,
		*crudeOil,
		*water,

		*leaves,
		*wood,
		*flowerPetals,
		*mycelia,
		*alienCarapace,
		*alienOrgans,
		*greenPowerSlug,
		*yellowPowerSlug,
		*purplePowerSlug,

		*ironIngot,
		*copperIngot,
		*steelIngot,
		*cateriumIngot,
		*aluminumIngot,

		*ironPlate,
		*ironRod,
		*screw,
		*reinforcedIronPlate,
		*modularFrame,
		*steelBeam,
		*steelPipe,
		*encasedIndustrialBeam,
		*heavyModularFrame,
		*alcladAluminumSheet,
		*copperSheet,

		*wire,
		*cable,
		*quickwire,
		*circuitBoard,
		*aiLimiter,
		*highSpeedConnector,
		*battery,

		*biomass,
		*solidBiofuel,
		*fabric,

		*concrete,
		*quartzCrystal,
		*silica,
		*compactedCoal,

		*rotor,
		*stator,
		*motor,
		*heatSink,
		*turboMotor,

		*powerShard,

		*plastic,
		*rubber,
		*petroleumCoke,
		*polymerResin,

		*fuel,
		*turbofuel,
		*liquidBiofuel,

		*packagedFuel,
		*packagedTurbofuel,
		*packagedLiquidBiofuel,
		*packagedHeavyOilResidue,
		*packagedOil,
		*packagedWater,

		*aluminumScrap,
		*aluminaSolution,
		*sulfuricAcid,
		*uraniumPellet,

		*emptyCanister,
		*heavyOilResidue,

		*smartPlating,
		*versatileFramework,
		*automatedWiring,
		*modularEngine,
		*adaptiveControlUnit,

		*crystalOscillator,
		*computer,
		*supercomputer,
		*radioControlUnit,

		*encasedUraniumCell,
		*electromagneticControlRod,
		*nuclearFuelRod,
	)

	return &Data{Items: &items}
}
