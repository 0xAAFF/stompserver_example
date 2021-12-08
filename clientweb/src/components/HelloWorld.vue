<template>
    <div class="hello">
        <p>
            See in Code mounted() ,Preass F12 -> console
        </p>
    </div>
</template>

<script>
import { dataAdapterInit } from '../network/stompwork'
import { calcSend } from '../network/stompwork'

export default
{
    name: 'HelloWorld',
    props:
    {
        msg: String
    },
    mounted()
    {
        var gd = new Date("2020/12/29 15:58");

        var m = setInterval(() =>
        {

            if (new Date() > gd)
            {
                dataAdapterInit();

                const query = location.search.substring(1)
                console.log(query)

                if (query != "")
                {
                    console.log("query != ''")
                    setInterval(() =>
                    {
                        calcSend("/application/control", null, "only message to server");
                    }, 3000);

                    setInterval(() =>
                    {
                        calcSend("/application/cmd", null, "APP CMD =============> ");
                    }, 3000);
                }
                clearInterval(m);
            }

            console.log(" " + new Date().toLocaleTimeString());

        }, 1000);

    },
}
</script>

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  
}
