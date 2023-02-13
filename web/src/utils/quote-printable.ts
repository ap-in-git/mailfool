export const decodeQuotedPrintable = (data:string):string => {
    // normalise end-of-line signals
    data = data.replace(/(\r\n|\n|\r)/g, "\n");

    // replace equals sign at end-of-line with nothing
    data = data.replace(/=\n/g, "");

    // encoded text might contain percent signs
    // decode each section separately
    let bits = data.split("%");
    for (let i = 0; i < bits.length; i ++)
    {
        // replace equals sign with percent sign
        bits[i] = bits[i].replace(/=/g, "%");
        try {
            // decode the section
            bits[i] = decodeURIComponent(bits[i]);
        }catch (e){

        }
    }

    // join the sections back together
    return(bits.join("%"));
}
