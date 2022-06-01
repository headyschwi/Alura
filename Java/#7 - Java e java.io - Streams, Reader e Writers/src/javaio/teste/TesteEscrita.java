package javaio.teste;

import java.io.IOException;
import java.io.Writer;
import java.io.OutputStream;
import java.io.OutputStreamWriter;
import java.io.BufferedWriter;
import java.io.FileOutputStream;

public class TesteEscrita {

	public static void main(String[] args) throws IOException {
		OutputStream fs = new FileOutputStream("lorem2.txt");
		Writer outWriter = new OutputStreamWriter(fs);
		BufferedWriter buffer = new BufferedWriter(outWriter);
		
		buffer.write("Hello, world!");
		buffer.newLine();
		buffer.write("Bye, world.");
		
		buffer.close();
	}

}
