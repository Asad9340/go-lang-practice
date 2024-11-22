package templates

templ Navbar() {
  <nav class="bg-blue-500 text-white py-4">
    <div class="container mx-auto flex justify-between">
      <h1 class="text-2xl font-bold">My Website</h1>
      <ul class="flex gap-4">
        <li><a href="/" class="hover:underline">Home</a></li>
        <li><a href="/about" class="hover:underline">About</a></li>
        <li><a href="/contact" class="hover:underline">Contact</a></li>
      </ul>
    </div>
  </nav>
}
