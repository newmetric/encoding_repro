package tx

import (
	"fmt"
	"testing"

	"cosmossdk.io/x/tx/signing"

	codecaddress "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/opinit/x/ophost"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/gogoproto/proto"

	cryptocodec "github.com/initia-labs/initia/crypto/codec"
)

func TestGoGoProtoJsonPB(t *testing.T) {
	interfaceRegistry, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          codecaddress.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: codecaddress.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		},
	})
	if err != nil {
		panic(err)
	}

	ophost.AppModuleBasic{}.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)

	appCodec := codec.NewProtoCodec(interfaceRegistry)
	decoder := DefaultJSONTxDecoder(appCodec)

	// const CHAIN_TYPE = "CHAIN_TYPE_CELESTIA" // success
	const CHAIN_TYPE = "CELESTIA" // failed (can't unmarshal Any nested proto *types.MsgCreateBridge: unknown value "CELESTIA" for enum opinit.ophost.v1.BatchInfo_ChainType)

	jsonText := fmt.Sprintf(`{"body":{"messages":[{"@type":"/opinit.ophost.v1.MsgCreateBridge","creator":"init1fgwvmvr8y5ul03aty9647vny7uj0uzt4g5zxev","config":{"challenger":"init1n8x7h4l96wmazlmxpqurccfg37fayrdafpmhfr","proposer":"init1knt8ehj03wk6hrzr4j35rmx2m7gqf4mwwcmrze","batch_info":{"submitter":"celestia13ycasdutemyk6pw6dy3ct3rxwknxrz6ygzjmu7","chain_type":"%s"},"submission_interval":"60s","finalization_period":"604800s","submission_start_height":"1","oracle_enabled":true,"metadata":"eyJwZXJtX2NoYW5uZWxzIjpbeyJwb3J0X2lkIjoibmZ0LXRyYW5zZmVyIiwiY2hhbm5lbF9pZCI6ImNoYW5uZWwtMzUzIn0seyJwb3J0X2lkIjoidHJhbnNmZXIiLCJjaGFubmVsX2lkIjoiY2hhbm5lbC0zNTIifV19"}}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AwKanlZ3AymXu7E54dOD5gL5kSrSBvTdrcow/vOsTVfl"},"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"sequence":"7"}],"fee":{"amount":[{"denom":"uinit","amount":"3000"}],"gas_limit":"200000","payer":"","granter":""},"tip":null},"signatures":["4x52q4ac4+Xbg/7PXzz73vhQrZgnL+FKsLuCYDK4f1EFOnsXvFAtyn+cqXwzXTJiKY+ZUoi4QfC4u+LEj0IZtg=="]}`, CHAIN_TYPE)
	_, err = decoder([]byte(jsonText))
	if err != nil {
		t.Fatal(err)
	}
}
