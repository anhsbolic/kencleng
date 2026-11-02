import { setupWorker } from 'msw/browser';
import { handlers, completeMockAuth, setScenario } from './handlers';
export const worker = setupWorker(...handlers);
export { completeMockAuth, setScenario };
