package javaio.teste;

import java.io.IOException;
import java.io.Writer;
import java.io.OutputStream;
import java.io.OutputStreamWriter;
import java.io.BufferedWriter;
import java.io.FileOutputStream;
import java.io.FileWriter;

public class TesteEscritaPrintStreamPrintWriter {

	public static void main(String[] args) throws IOException {
		//OutputStream fs = new FileOutputStream("lorem2.txt");
		//Writer outWriter = new OutputStreamWriter(fs);
		//BufferedWriter buffer = new BufferedWriter(outWriter);
		
		BufferedWriter bw = new BufferedWriter(new FileWriter("lorem2.txt"));
		
		bw.write("Hello, world!");
		bw.newLine();
		bw.write("Bye, world.");
		
		bw.close();
	}

}
