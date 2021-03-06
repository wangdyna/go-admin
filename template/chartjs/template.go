package chartjs

var List = map[string]string{
	"chartjs": `{{define "chartjs"}}
    {{if ne .Title ""}}
        <p class="text-center">
            <strong>{{langHtml .Title}}</strong>
        </p>
    {{end}}
    <div class="chart">
        <canvas id="{{.ID}}" style="height: {{.Height}}px;"></canvas>
    </div>
    <script>
        new Chart(document.getElementById('{{.ID}}'), {{.Js}});
    </script>
{{end}}`,
}
