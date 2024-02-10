import { ZenEngine } from '@gorules/zen-engine';
import fs from 'fs/promises';

const engine = new ZenEngine();

const evaluate = async () => {
    const lol = performance.now();
    const content = await fs.readFile('./test.json');
    const decision = engine.createDecision(content);
    const result = await decision.evaluate({
		"MonthlyCashInflow": 10,
	});
    const lol2 = performance.now();
    console.log(lol2 - lol);
    console.log(result)
};

evaluate();
