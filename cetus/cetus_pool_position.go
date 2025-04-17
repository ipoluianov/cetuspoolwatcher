package cetus

/*
{
  "dataType": "moveObject",
  "fields": {
    "coin_type_a": {
      "fields": {
        "name": "356a26eb9e012a68958082340d4c4116e7f55615cf27affcff209cf0ae544f59::wal::WAL"
      },
      "type": "0x1::type_name::TypeName"
    },
    "coin_type_b": {
      "fields": {
        "name": "0000000000000000000000000000000000000000000000000000000000000002::sui::SUI"
      },
      "type": "0x1::type_name::TypeName"
    },
    "description": "Cetus Liquidity Position",
    "id": {
      "id": "0x99e5651cc2022b4cd688da8fd7cb250efe545ae1d4984539c0f484982b86aeb6"
    },
    "index": "8",
    "liquidity": "896946922429",
    "name": "Cetus LP | Pool26051-8",
    "pool": "0xf4238fa592c9ed7f148fd091cb2c4147cb15ad81b797115ce42971923ebf6e4c",
    "tick_lower_index": {
      "fields": {
        "bits": 4294950176
      },
      "type": "0x714a63a0dba6da4f017b42d5d0fb78867f18bcde904868e51d951a5a6f5b7f57::i32::I32"
    },
    "tick_upper_index": {
      "fields": {
        "bits": 4294952116
      },
      "type": "0x714a63a0dba6da4f017b42d5d0fb78867f18bcde904868e51d951a5a6f5b7f57::i32::I32"
    },
    "url": "https://bq7bkvdje7gvgmv66hrxdy7wx5h5ggtrrnmt66rdkkehb64rvz3q.arweave.net/DD4VVGknzVMyvvHjceP2v0_TGnGLWT96I1KIcPuRrnc"
  },
  "hasPublicTransfer": true,
  "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::position::Position"
}

*/

type CetusPoolPosition struct {
	DataType string `json:"dataType"`
	Fields   struct {
		CoinTypeA struct {
			Fields struct {
				Name string `json:"name"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"coin_type_a"`
		CoinTypeB struct {
			Fields struct {
				Name string `json:"name"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"coin_type_b"`
		Description string `json:"description"`
		Id          struct {
			Id string `json:"id"`
		} `json:"id"`
		Index          string `json:"index"`
		Liquidity      string `json:"liquidity"`
		Name           string `json:"name"`
		Pool           string `json:"pool"`
		TickLowerIndex struct {
			Fields struct {
				Bits int `json:"bits"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"tick_lower_index"`
		TickUpperIndex struct {
			Fields struct {
				Bits int `json:"bits"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"tick_upper_index"`
		Url string `json:"url"`
	} `json:"fields"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Type              string `json:"type"`
}
