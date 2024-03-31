{{template "top" .}}

<div class="posts">
  <div class="width">
    <h1>Posts</h1>

    <div class="post-table">
      <div class="post-row head">
        <div class="date">Date</div>
        <div class="price">Price</div>
        <div class="title">Title</div>
        <div class="content">Content</div>
        <div class="username">username</div>
      </div>

      {{ range.posts }}
      <div class="post-row">
        <div class="date">
          {{ .CreateAt }}
        </div>
        <div class="price">
          {{ .Price }}
        </div>
        <div class="title">
          {{ .Title }}
        <div class="content">
          {{ .Content }}
        </div>
        <div class="username">
          {{ .Username }}
        </div>
      </div>
      {{ end }}
    </div>
  </div>
</div>

{{template "bottom" .}}