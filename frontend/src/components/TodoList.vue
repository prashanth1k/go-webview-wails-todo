<template>
    <div class="todolist">
        <div v-for="todo in todos.data" :key="todo">
            <div class="todo-item row">
                <div class="col-1">
                    <input type="checkbox" name="status" id="status" :checked="todo.status == 'Complete' ? true : false"
                        @click="updateTodo(todo)">
                </div>
                <div class="col-3 dates">
                    <img src="https://icongr.am/feather/calendar.svg?size=14&color=D3D3D3" alt=""> {{
                        todo.due?.substring(0, 10) }}
                </div>
                <div class="col-7">
                    {{ todo.description }}
                </div>
                <div class="col-1">
                    <a class="button error cardaction " @click="deleteTodo(todo)"><img
                            src="https://icongr.am/feather/trash-2.svg?size=24&color=FFFFFF"></a>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
const props = defineProps(['todos'])
import { UpdateTodo, DeleteTodo } from '../../wailsjs/go/main/App'
const emit = defineEmits(['updateTodo', 'deleteTodo'])

const updateTodo = async (todo) => {
    console.log(`Emitting updateTodo event..!`)
    todo.status = todo.status == "Complete" ? "In Progress" : "Complete"

    const todoResponse = await UpdateTodo(todo)
    console.log('todoResponse: ', todoResponse);
    emit('updateTodo', todo)
}

const deleteTodo = async (todo) => {
    console.log(`Emitting deleteTodo event..!`)
    const todoResponse = await DeleteTodo(todo)
    console.log('todoResponse: ', todoResponse);
    emit('deleteTodo', todo)
}

</script>
