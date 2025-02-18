package rpc

import (
	"github.com/sev-2/raiden"
	"beraiden/internal/models"
)

type CustomAccessTokenHookParams struct {
	Event map[string]interface{} `json:"event" column:"name:event;type:jsonb"`
}
type CustomAccessTokenHookResult map[string]interface{}

type CustomAccessTokenHook struct {
	raiden.RpcBase
	Params   *CustomAccessTokenHookParams `json:"-"`
	Return   CustomAccessTokenHookResult `json:"-"`
}

func (r *CustomAccessTokenHook) GetName() string {
	return "custom_access_token_hook"
}

func  (r *CustomAccessTokenHook) UseParamPrefix() bool {
	return false
}

func (r *CustomAccessTokenHook) GetBehavior() raiden.RpcBehaviorType {
	return raiden.RpcBehaviorStable
}

func (r *CustomAccessTokenHook) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeJSONB
}

func (r *CustomAccessTokenHook) BindModels() {
	r.BindModel(models.UserRoles{}, "u")
}

func (r *CustomAccessTokenHook) GetRawDefinition() string {
	return `declare claims jsonb; user_role public.app_role; begin select role into user_role from public.:u where user_id = (:event->>'user_id')::uuid; claims := :event->'claims'; if user_role is not null then claims := jsonb_set(claims, '{user_role}', to_jsonb(user_role)); else claims := jsonb_set(claims, '{user_role}', 'null'); end if; :event := jsonb_set(:event, '{claims}', claims); return :event; end;`
}