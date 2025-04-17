package cetus

/*
{
  "dataType": "moveObject",
  "fields": {
    "coin_a": "7444207322610085",
    "coin_b": "451508345018186",
    "current_sqrt_price": "8114169058997184263",
    "current_tick_index": {
      "fields": {
        "bits": 4294950869
      },
      "type": "0x714a63a0dba6da4f017b42d5d0fb78867f18bcde904868e51d951a5a6f5b7f57::i32::I32"
    },
    "fee_growth_global_a": "2508122698513772346",
    "fee_growth_global_b": "429590139909167098",
    "fee_protocol_coin_a": "10794277015952",
    "fee_protocol_coin_b": "1736555226720",
    "fee_rate": "2500",
    "id": {
      "id": "0x72f5c6eef73d77de271886219a2543e7c29a33de19a6c69c5cf1899f729c3f17"
    },
    "index": "25952",
    "is_pause": false,
    "liquidity": "190398899131800498",
    "position_manager": {
      "fields": {
        "position_index": "205438",
        "positions": {
          "fields": {
            "head": "0xbab6e72ba509f5cd1d1290aebfe58a464e2819a3505609d1156e69761ae97f60",
            "id": {
              "id": "0xa699b6c9380b07f5077ac5c5471ac5e0edffc409b8d94469c3cb6bf37671ad5a"
            },
            "size": "1258",
            "tail": "0x07d8ef9ad971ee5ecc785daa4bcfe17918397836a2f17b75ccd9055eb9c7c508"
          },
          "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::linked_table::LinkedTable\u003c0x2::object::ID, 0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::position::PositionInfo\u003e"
        },
        "tick_spacing": 60
      },
      "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::position::PositionManager"
    },
    "rewarder_manager": {
      "fields": {
        "last_updated_time": "1744901557",
        "points_growth_global": "1337496053459358",
        "points_released": "33791409877007810194046976000000",
        "rewarders": [
          {
            "fields": {
              "emissions_per_second": "425726940775194975952592593",
              "growth_global": "15184226909956425",
              "reward_coin": {
                "fields": {
                  "name": "0000000000000000000000000000000000000000000000000000000000000002::sui::SUI"
                },
                "type": "0x1::type_name::TypeName"
              }
            },
            "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::rewarder::Rewarder"
          },
          {
            "fields": {
              "emissions_per_second": "11496335432788941161054814815",
              "growth_global": "407189968125860411",
              "reward_coin": {
                "fields": {
                  "name": "356a26eb9e012a68958082340d4c4116e7f55615cf27affcff209cf0ae544f59::wal::WAL"
                },
                "type": "0x1::type_name::TypeName"
              }
            },
            "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::rewarder::Rewarder"
          }
        ]
      },
      "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::rewarder::RewarderManager"
    },
    "tick_manager": {
      "fields": {
        "tick_spacing": 60,
        "ticks": {
          "fields": {
            "head": [
              {
                "fields": {
                  "is_none": false,
                  "v": "56"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "333236"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "333236"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "333236"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "333236"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "411476"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "435056"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "435056"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "435056"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              },
              {
                "fields": {
                  "is_none": false,
                  "v": "435056"
                },
                "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
              }
            ],
            "id": {
              "id": "0x87a742dfd5ea41a9e995dc90b7acc3d96dc0442bf3cf0c6ff9c5d90268b50757"
            },
            "level": "10",
            "list_p": "2",
            "max_level": "16",
            "random": {
              "fields": {
                "seed": "2350102922056326544"
              },
              "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::random::Random"
            },
            "size": "358",
            "tail": {
              "fields": {
                "is_none": false,
                "v": "887216"
              },
              "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::option_u64::OptionU64"
            }
          },
          "type": "0xbe21a06129308e0495431d12286127897aff07a8ade3970495a4404d97f9eaaa::skip_list::SkipList\u003c0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::tick::Tick\u003e"
        }
      },
      "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::tick::TickManager"
    },
    "tick_spacing": 60,
    "url": "https://bq7bkvdje7gvgmv66hrxdy7wx5h5ggtrrnmt66rdkkehb64rvz3q.arweave.net/DD4VVGknzVMyvvHjceP2v0_TGnGLWT96I1KIcPuRrnc"
  },
  "hasPublicTransfer": true,
  "type": "0x1eabed72c53feb3805120a081dc15963c204dc8d091542592abaf7a35689b2fb::pool::Pool\u003c0x356a26eb9e012a68958082340d4c4116e7f55615cf27affcff209cf0ae544f59::wal::WAL, 0x2::sui::SUI\u003e"
}*/

type CetusPool struct {
	DataType string `json:"dataType"`
	Fields   struct {
		CoinA            string `json:"coin_a"`
		CoinB            string `json:"coin_b"`
		CurrentSqrtPrice string `json:"current_sqrt_price"`
		CurrentTickIndex struct {
			Fields struct {
				Bits int `json:"bits"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"current_tick_index"`
		FeeGrowthGlobalA string `json:"fee_growth_global_a"`
		FeeGrowthGlobalB string `json:"fee_growth_global_b"`
		FeeProtocolCoinA string `json:"fee_protocol_coin_a"`
		FeeProtocolCoinB string `json:"fee_protocol_coin_b"`
		FeeRate          string `json:"fee_rate"`
		Id               struct {
			Id string `json:"id"`
		} `json:"id"`
		Index           string `json:"index"`
		IsPause         bool   `json:"is_pause"`
		Liquidity       string `json:"liquidity"`
		PositionManager struct {
			Fields struct {
				PositionIndex string `json:"position_index"`
				Positions     struct {
					Fields struct {
						Head string `json:"head"`
						Id   struct {
							Id string `json:"id"`
						} `json:"id"`
						Size string `json:"size"`
						Tail string `json:"tail"`
					} `json:"fields"`
					Type string `json:"type"`
				} `json:"positions"`
				TickSpacing string `json:"tick_spacing"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"position_manager"`
		RewarderManager struct {
			Fields struct {
				LastUpdatedTime    string `json:"last_updated_time"`
				PointsGrowthGlobal string `json:"points_growth_global"`
				PointsReleased     string `json:"points_released"`
				Rewarders          []struct {
					Fields struct {
						EmissionsPerSecond string `json:"emissions_per_second"`
						GrowthGlobal       string `json:"growth_global"`
						RewardCoin         struct {
							Fields struct {
								Name string `json:"name"`
							} `json:"fields"`
							Type string `json:"type"`
						} `json:"reward_coin"`
					} `json:"fields"`
					Type string `json:"type"`
				} `json:"rewarders"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"rewarder_manager"`
		TickManager struct {
			Fields struct {
				TickSpacing string `json:"tick_spacing"`
				Ticks       struct {
					Fields struct {
						Head []struct {
							Fields struct {
								IsNone bool   `json:"is_none"`
								V      string `json:"v"`
							} `json:"fields"`
							Type string `json:"type"`
						} `json:"head"`
						Id struct {
							Id string `json:"id"`
						} `json:"id"`
						Level    string `json:"level"`
						ListP    string `json:"list_p"`
						MaxLevel string `json:"max_level"`
						Random   struct {
							Fields struct {
								Seed string `json:"seed"`
							} `json:"fields"`
							Type string `json:"type"`
						} `json:"random"`
						Size string `json:"size"`
						Tail struct {
							Fields struct {
								IsNone bool   `json:"is_none"`
								V      string `json:"v"`
							} `json:"fields"`
							Type string `json:"type"`
						} `json:"tail"`
					} `json:"fields"`
					Type string `json:"type"`
				} `json:"ticks"`
			} `json:"fields"`
			Type string `json:"type"`
		} `json:"tick_manager"`
		TickSpacing string `json:"tick_spacing"`
		Url         string `json:"url"`
	} `json:"fields"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Type              string `json:"type"`
}
