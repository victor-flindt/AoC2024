Task 1:
Scan corrupted memory for valid multiplication instructions (mul(X,Y)), ignoring any invalid characters or sequences. For each valid multiplication instruction, compute the result and sum all the results.

Task 2:
In addition to the valid mul(X,Y) instructions, handle the new instructions do() and don't(), which enable or disable multiplication instructions, respectively. Only compute and sum the results of mul instructions that are enabled based on the most recent do() or don't() instruction.