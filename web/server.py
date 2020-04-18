from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def index():
    return "It works."

@app.route("/live")
def live():
    return render_template("index.html")

if __name__ == "__main__":
    app.run("localhost", 9999)