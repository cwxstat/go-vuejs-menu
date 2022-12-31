<template>
  <div class="detail">
    

    <h3>
      References:
    </h3>

    <a href="https://www.youtube.com/watch?v=H-hXNym2CK8">video 3</a>
    <br />
    <a href="https://youtu.be/GdWrYfDfqRE?t=6">video 4</a>
    <br />
    <a href="https://youtu.be/usSBsgWNUZk?t=9">video 5</a>

    <div class="AppDev">



      {{ name }}, {{ age }}
      <br /><br />
      <button @click="changeName('mike')">Change Name</button>
      <button @click="grabData()">GrabData</button>
      <br /> <br /><br />

      <div :style="{ color: activeColor, fontSize: fontSize + 'px' }">
        {{ msg }}
        <br /> <br />
        Ref: <a
          href="https://www.typescriptlang.org/play?#code/JYOwLgpgTgZghgYwgAgAoEEpTgT2QbwFgAoZM5ABwC5kQBXAWwCNoBuEkgXw+NElkQoAInDBwA8kwBWBEuWQATUXBoYsuANoBddqXJgcFCDQDOYKKADmu7sRIAbCGGQB7aTRFjJMgLzIAUgDK4gByAHQUcFAmEAAUAOT4AERKYkk0GskU6cgALACMYQCs+ZwANARJ2TQAzIUATAAMnFoVSQZGOe3YChBJnPEAlLoILiAmLo5h9i6WsW5SYalwGo1aEcMko+OTENOz89JhHRDDQA">TypeScript</a>
        <br />
        {{ obj.data[0].p }}
      </div>

      <br />
      <button @click="sendMessage('Hello websocket')">Send Message</button>
      <br />
      <button @click="stopMessage()">Stop Message</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import axios from "axios";
import { reactive } from 'vue'

interface PArray {
  p: number;
}

interface DataObj {
  data: PArray[];
  type: string;
}


export default defineComponent({
  name: 'AppDev',
  components: {},
  setup() {

    const activeColor = ref('red')
    const fontSize = ref(15)

    // Define your own type.. DataObj
    const obj = ref<DataObj>({ "data": [{ "p": 42.05 }, { "p": 40.68 }], "type": "buy" })

    const name = ref('Susan');
    const msg = ref('msg');
    const age = ref<number | string>(60)
    const connection = ref<any>({});
    return { name, age, connection, msg, activeColor, fontSize, obj }
  },
  beforeRouteLeave(to, from, next) {
    console.log('beforeRouteLeave... close websocket');
    
    this.connection.onclose = function () {
      console.log('WebSocket Client Closed');
    };
    this.connection.close();
    next();
  },
  created() {
    console.log("Starting Connection to WebSocket Server")
    // MAKE SURE TO CHANGE THE URL TO YOUR LOCAL IP ADDRESS
    // 1 of 2 (Search for 2 of 2)
    // if on a permanent location. 
    // this.connection = new WebSocket('wss://cwxstat.com/ws');
    // Uncomment below... testing or local
    this.connection = new WebSocket('ws://localhost:8080/ws');

    this.connection.onopen = (event: string) => {
      console.log("Connection to WebSocket Server Established")
      console.log(event)
    }
    this.connection.onmessage = (event: any) => {
      console.log("Message Received")
      this.msg = event.data
      this.updateObj(event.data)
      console.log(event)
    }
  },
  methods: {
    updateObj(j: string) {
      this.obj = JSON.parse(j)
      // Change color
      if (this.obj.type == "buy") {
        this.activeColor = "green"
      } else {
        this.activeColor = "red"
      }

    },
    changeName(name: string) {
      this.name = name
      return name
    },
    changeAge(age: string | number) {
      this.age = age
      return age
    },
    grabData() {
      const article = { title: "Vue POST Request Example" };
      axios.post("https://reqres.in/api/articles", article)
        .then(response => this.age = response.data.id)
        .catch(error => {
          console.log(error)
        }
        )
    },
    reConnect() {
      console.log("Starting Connection to WebSocket Server")
      // 2 of 2
      // this.connection = new WebSocket('wss://cwxstat.com/ws');
      // CHANGE if on a permanent location.
      this.connection = new WebSocket('ws://localhost:8080/ws');
      this.connection.onopen = (event: any) => {
        console.log("Connection to WebSocket Server Established")
        console.log(event)
      }
      this.connection.onmessage = (event: any) => {
        console.log("Message Received")
        this.msg = event.data
        this.updateObj(event.data)
        console.log(event)
      }

    },
    sendMessage(message: string) {
      this.connection.send(message)
    },
    stopMessage() {
      this.connection.close()
      this.msg = "Connection Closed"
      this.reConnect()
    },
  },

});

</script>


<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>