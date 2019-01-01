<template>
  <ul>
    <li><input
      class="todo-form"
      placeholder="What needs to be done?"
      @keyup.enter="addTodo"></li>
    <li
      v-for="(todo, index) in todos"
      :key="index">
      <input
        :checked="todo.done"
        type="checkbox"
        @change="toggle(todo)">
      <span :class="{ done: todo.done }">{{ todo.text }}</span>
      <input
        class="remove-btn"
        type="button"
        value="remove"
        @click="remove(todo)">
    </li>
  </ul>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  computed: {
    todos() {
      return this.$store.state.todos.list
    }
  },
  methods: {
    addTodo(e) {
      this.$store.commit('todos/add', e.target.value)
      e.target.value = ''
    },
    ...mapMutations({
      remove: 'todos/remove',
      toggle: 'todos/toggle'
    })
  }
}
</script>

<style>
li {
  list-style-type: none;
}

.done {
  text-decoration: line-through;
}

.todo-form {
  border: 0;
  padding: 10px;
  font-size: 1.3em;
  font-family: Arial, sans-serif;
  color: #aaa;
  border: solid 1px #ccc;
  margin: 0 0 20px;
  width: 300px;
}

.remove-btn {
  margin-left: 20px;
}
</style>
