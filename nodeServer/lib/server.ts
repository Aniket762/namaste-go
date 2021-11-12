import express, { Application, Request, Response } from "express";

const app: Application = express();
const PORT: number = 8080;

// middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (_, res: Response) => {
  res.send("<h1>Welcome to Go</h1>");
});

app.get("/get", (_, res: Response) => {
  res.status(200).json({ message: "Hello from Aniket" });
});

app.post("/post", (req: Request, res: Response) => {
  let gJSON = req.body;
  res.status(200).send(JSON.stringify(req.body));
});

app.post("/postform", (req: Request, res: Response) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => console.log("Server running at port: " + PORT));
