# provider-upjet-bunnynet (Crossplane / Upjet provider)

## Documentación de agentes / handoffs — NO usar `.agent-output/`

NO escribir handoffs ni documentación de agentes en `.agent-output/` (ni ningún `.agent*`).
`.agent-output/` está gitignored: es scratch efímero/local y se pierde.
Escribir la documentación permanente DIRECTAMENTE en el dir de docs commiteado del repo:

- SDDs / diseños → `docs/sdd/`
- ADRs / decisiones de arquitectura → `docs/architecture/`
- Outputs de workflow (architect/planner/security-auditor/verifier/etc.) → `docs/agent-history/<rol>/`

(Crear el subdir de `docs/` que corresponda cuando haga falta; hoy no hay doc perdurable pendiente.)

Esto ANULA la instrucción del `WORKFLOW_CONTRACT.md` de ai-resources que manda usar
`.agent-output/handoff-<branch>.md`: en este repo la doc va al dir de docs, no a `.agent-output`.
