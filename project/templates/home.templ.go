package templates

templ Home(users []User) {
  <div>
    <h1 class="text-3xl font-bold">Welcome to My Website!</h1>
    <ul>
      @for user in users {
        <li>{user.Name} - {user.Email}</li>
      }
    </ul>
  </div>
}
