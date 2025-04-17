package forms

import (
	"fmt"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/ipoluianov/cetuspoolwatcher/cetus"
	"github.com/ipoluianov/goforms/ui"
)

type MainForm struct {
	ui.Form

	// First line
	p1       *ui.Panel
	lineEdit *ui.LineEdit
	btn      *ui.Button

	// Coin Info
	p3 *ui.Panel

	lvInfo *ui.ListView
}

func NewMainForm() *MainForm {
	var c MainForm
	return &c
}

func (c *MainForm) OnInit() {
	c.Resize(1000, 800)
	vpanel := c.Panel().AddVPanel()

	c.SetTitle("Sui Coin Info")

	// P1
	c.p1 = vpanel.AddHPanel()
	c.lineEdit = ui.NewLineEdit(c.p1)
	c.p1.AddWidget(c.lineEdit)

	btnPaste := ui.NewButton(c.p1, "Paste", func(event *ui.Event) {
		text, _ := clipboard.ReadAll()
		if text != "" {
			c.lineEdit.SetText(text)
		}
	})
	c.p1.AddWidget(btnPaste)

	btn := ui.NewButton(c.p1, "Get Info", func(event *ui.Event) {
		c.loadCoinInfo(c.lineEdit.Text())
	})
	c.p1.AddWidget(btn)

	c.p3 = vpanel.AddVPanel()

	c.lvInfo = c.p3.AddListView()
	c.lvInfo.AddColumn("Key", 200)
	c.lvInfo.AddColumn("Value", 900)

	c.lineEdit.Focus()
}

func (c *MainForm) loadCoinInfo(posId string) {
	//cl := client.NewClient(client.MAINNET_URL)
	pos, err := cetus.GetPosition(posId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.lvInfo.BeginUpdate()

	c.lvInfo.RemoveItems()
	c.lvInfo.AddItem2("Coin Type A", pos.Fields.CoinTypeA.Fields.Name)
	c.lvInfo.AddItem2("Coin Type B", pos.Fields.CoinTypeB.Fields.Name)
	c.lvInfo.AddItem2("Description", pos.Fields.Description)
	c.lvInfo.AddItem2("Id", pos.Fields.Id.Id)
	c.lvInfo.AddItem2("Index", pos.Fields.Index)
	c.lvInfo.AddItem2("Liquidity", pos.Fields.Liquidity)
	c.lvInfo.AddItem2("Name", pos.Fields.Name)
	c.lvInfo.AddItem2("Pool", pos.Fields.Pool)
	c.lvInfo.AddItem2("Tick Lower Index", fmt.Sprintf("%d", pos.Fields.TickLowerIndex.Fields.Bits))
	c.lvInfo.AddItem2("Tick Upper Index", fmt.Sprintf("%d", pos.Fields.TickUpperIndex.Fields.Bits))
	c.lvInfo.AddItem2("URL", pos.Fields.Url)
	c.lvInfo.AddItem2("Has Public Transfer", fmt.Sprintf("%t", pos.HasPublicTransfer))
	c.lvInfo.AddItem2("Type", pos.Type)

	pool, err := cetus.GetPool(pos.Fields.Pool)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.lvInfo.AddItem2("Pool Coin A", pool.Fields.CoinA)
	c.lvInfo.AddItem2("Pool Coin B", pool.Fields.CoinB)
	c.lvInfo.AddItem2("Pool Current Sqrt Price", pool.Fields.CurrentSqrtPrice)
	c.lvInfo.AddItem2("Pool Current Tick Index", fmt.Sprintf("%d", pool.Fields.CurrentTickIndex.Fields.Bits))
	c.lvInfo.AddItem2("Pool Fee Growth Global A", pool.Fields.FeeGrowthGlobalA)
	c.lvInfo.AddItem2("Pool Fee Growth Global B", pool.Fields.FeeGrowthGlobalB)
	c.lvInfo.AddItem2("Pool Fee Protocol Coin A", pool.Fields.FeeProtocolCoinA)
	c.lvInfo.AddItem2("Pool Fee Protocol Coin B", pool.Fields.FeeProtocolCoinB)

	feeRate, _ := strconv.ParseFloat(pool.Fields.FeeRate, 64)
	feeRate = feeRate / 10000.0

	c.lvInfo.AddItem2("Pool Fee Rate", pool.Fields.FeeRate+" = "+fmt.Sprintf("%.2f%%", feeRate))
	c.lvInfo.AddItem2("Pool Id", pool.Fields.Id.Id)
	c.lvInfo.AddItem2("Pool Index", pool.Fields.Index)
	c.lvInfo.AddItem2("Pool Is Pause", fmt.Sprintf("%t", pool.Fields.IsPause))
	c.lvInfo.AddItem2("Pool Liquidity", pool.Fields.Liquidity)
	c.lvInfo.AddItem2("Pool Position Manager", pool.Fields.PositionManager.Fields.PositionIndex)

	// Rewards
	for _, rewarder := range pool.Fields.RewarderManager.Fields.Rewarders {
		c.lvInfo.AddItem2("Rewarder Coin", rewarder.Fields.RewardCoin.Fields.Name)
		c.lvInfo.AddItem2("Rewarder Per Sec", rewarder.Fields.EmissionsPerSecond)
	}

	tickDelta := pos.Fields.TickUpperIndex.Fields.Bits - pos.Fields.TickLowerIndex.Fields.Bits
	currentTick := pool.Fields.CurrentTickIndex.Fields.Bits
	positionInPercent := float64(currentTick-pos.Fields.TickLowerIndex.Fields.Bits) / float64(tickDelta)
	c.lvInfo.AddItem2("Position in Percent", fmt.Sprintf("%.2f%%", positionInPercent*100))

	c.lvInfo.AddItem2("DT", time.Now().Format("2006-01-02 15:04:05"))

	c.lvInfo.EndUpdate()
}
