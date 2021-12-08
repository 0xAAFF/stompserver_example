export const projectconfig =
{
    stompHost: "ws://"+ window.location.host+"/stomp",    // 数据适配器地址
    SubscribeMap:
    {
        basicInformation: "/topic/suggestion",
        cmdInformation: "/topic/cmd",
        
    },

    sendMap:
    {
        send_cmd : "/app/cmd",
    }
};