void mainImage( out vec4 fragColor, in vec2 fragCoord )
{
    vec2 coord = fragCoord - (iResolution.xy / vec2(2.0));
   	coord *= sqrt(iTime);
    float x = float(coord.x);
    float y = float(coord.y);
    float r2 = float(x*x + y*y);      ;
    int a = int(floor(r2*pow(16.0, 6.0-ceil(log2(r2)/4.0))));   
    fragColor = vec4(
        float((a >> 16) & 255) / 255.0,
        float((a >> 8) & 255) / 255.0,
        float((a >> 0) & 255) / 255.0,
        1.0
    );
}
