<!doctype html>
<html lang="en">

<head>
  <title>Golang Todo App</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <script type="text/javascript" src="https://unpkg.com/vue@2.3.4"></script>
  <script src="https://cdn.jsdelivr.net/npm/vue-resource@1.3.4"></script>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css"
    integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
  <link rel="stylesheet" href="../static/style/style.css">
</head>
<style>
  .card {
    border-radius: 0;
    border: none;
  }

  .card-body {
    margin-top: 5px;
    padding: 0;
  }

  .todo-input {
    border-radius: 0;
    padding: 10px 10px;
    border-bottom: none;
  }

  .todo-input:focus,
  .todo-input:active {
    box-shadow: none;
  }

  .action-btn {
    margin-left: 5px;
    border-radius: 5px;
    cursor: pointer;
  }

  .action-btn:hover {
    background-color: #146b43;
  }

  .action-btn:focus,
  .action-btn:active {
    box-shadow: none;
  }

  .list-group li {
    cursor: pointer;
    border-radius: 5px;
  }

  .checked {
    background: #448197;
    color: #c0d1cf;
  }

  .show-error {
    color: red;
    display: block;
  }

  .del {
    text-decoration: line-through;
  }

  .not-checked {
    background: #67b48d;
    color: #FFF;
    font-weight: bold;
  }

  .title {
    width: 100%;
    padding: 8px;
    color: #15063f;
    background-color: #67b48d;
    text-align: center;
    font-size: 20px;
    font-weight: 700;
    border-radius: 2px;
  }
</style>

<body>
  <div class="container" id="app">
    <div class="row">
      <div class="col-12 mt-5 title">
        Goland Todo App
      </div>
      <div class="col-6 left-card">
        <div class="card">
          <div class="card-body">
            <form v-on:submit.prevent>
              <div class="input-group">
                <input type="text" v-model="todo.title" v-on:keyup="checkForEnter($event)"
                  class="form-control todo-input" placeholder="Write Todo here">
                <span class="input-group-btn">
                  <button class="btn custom-button" :class="{'btn-success' : !enableEdit, 'btn-warning' : enableEdit}"
                    type="button" v-on:click="addTodo"><span
                      :class="{'fa fa-plus' : !enableEdit, 'fa fa-edit' : enableEdit}"></span></button>
                </span>
              </div>
              <div class="show-error" v-if="showError">
                Field cannot be Blank
              </div>
            </form>
          </div>
        </div>
      </div>
      <div class="col-6 right-card">
        <div class="card"></div>
        <div class="card-body">
          <ul class="list-group">
            <li class="list-group-item" :class="{ 'checked': todo.completed, 'not-checked': !todo.completed }"
              v-for="(todo, todoIndex) in todos" v-on:click="toggleTodo(todo, todoIndex)">
              <i
                :class="{'fa fa-circle': !todo.completed, 'fa fa-check-circle text-success': todo.completed }">&nbsp;</i>
              <span :class="{ 'del': todo.completed }">@{ todo.title }</span>
              <div class="btn-group float-right" role="group" aria-label="Basic example">
                <button type="button" class="btn btn-success btn-sm action-btn m-1" v-on:click.prevent.stop
                  v-on:click="editTodo(todo, todoIndex)">Edit</button>
                <button type="button" class="btn btn-danger btn-sm action-btn m-1" v-on:click.prevent.stop
                  v-on:click="deleteTodo(todo, todoIndex)">Delete</button>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  </div>
  <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"
    integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
    crossorigin="anonymous"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js"
    integrity="sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh"
    crossorigin="anonymous"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js"
    integrity="sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ"
    crossorigin="anonymous"></script>
  <script type="text/javascript">
    var Vue = new Vue({
      el: '#app',
      delimiters: ['@{', '}'],
      data: {
        showError: false,
        enableEdit: false,
        todo: { id: '', title: '', completed: false },
        todos: []
      },
      mounted() {
        console.log("mopunted on")
        this.$http.get('getalltodos').then(response => {
          this.todos = response.body;
          console.log("response data", response)
        });
      },
      methods: {
        addTodo() {
          if (this.todo.title == '') {
            this.showError = true;
          } else {
            this.showError = false;
            if (this.enableEdit) {
              this.$http.put('updatetodo?id=' + this.todo.id, this.todo).then(response => {
                if (response.status == 200) {
                  this.todos[this.todo.todoIndex] = this.todo;
                }
              });
              this.todo = { id: '', title: '', completed: false };
              this.enableEdit = false;
            } else {
              this.$http.post('todo', { title: this.todo.title }).then(response => {
                if (response.status == 201) {
                  this.todos.push({ id: response.body.todo_id, title: this.todo.title, completed: false });
                  this.todo = { id: '', title: '', completed: false };
                }
              });
            }
          }
        },
        checkForEnter(event) {
          if (event.key == "Enter") {
            this.addTodo();
          }
        },
        toggleTodo(todo, todoIndex) {
          var completedToggle;
          if (todo.completed == true) {
            completedToggle = false;
          } else {
            completedToggle = true;
          }
          this.$http.put('updatetodo?id=' + todo.id, { id: todo.id, title: todo.title, completed: completedToggle }).then(response => {
            if (response.status == 200) {
              this.todos[todoIndex].completed = completedToggle;
              location.reload();
            }
          });
        },
        editTodo(todo, todoIndex) {
          this.enableEdit = true;
          this.todo = todo;
          this.todo.todoIndex = todoIndex;
        },
        deleteTodo(todo, todoIndex) {
          if (confirm("Are you sure ?")) {
            this.$http.delete('deletetodo?id=' + todo.id).then(response => {
              if (response.status == 200) {
                this.todos.splice(todoIndex, 1);
                this.todo = { id: '', title: '', completed: false };
              }
            });
          }
        }
      }
    });
  </script>
</body>

</html>