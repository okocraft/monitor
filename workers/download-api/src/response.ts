export function renderGZip(body: ReadableStream) {
	return new Response(body, {
		headers: {
			"Content-Type": "application/json",
			"Content-Encoding": "gzip",
		},
	});
}

export function errorResponse(
	statusCode: number,
	log: string,
	printError: boolean,
) {
	if (printError) {
		500 <= statusCode ? console.error(log) : console.warn(log);
	}
	return new Response(null, { status: statusCode });
}
