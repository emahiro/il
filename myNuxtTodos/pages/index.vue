<template>
  <section class="container">
    <div>
      <h1>This is MyTodo App</h1>
      <todos/>
    </div>
  </section>
</template>

<script>
import Logo from '~/components/Logo.vue'
import Todos from '~/components/Todos.vue'

export default {
  components: {
    Logo,
    Todos
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}

.todos {
  text-align: center;
}
</style>
