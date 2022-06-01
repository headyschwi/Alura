package javaio.teste;

import java.io.IOException;
import java.io.Writer;
import java.io.OutputStream;
import java.io.OutputStreamWriter;
import java.io.PrintStream;
import java.io.PrintWriter;
import java.io.BufferedWriter;
import java.io.FileOutputStream;
import java.io.FileWriter;

public class TesteEscritaFileWriter {

	public static void main(String[] args) throws IOException {
		//OutputStream fs = new FileOutputStream("lorem2.txt");
		//Writer outWriter = new OutputStreamWriter(fs);
		//BufferedWriter buffer = new BufferedWriter(outWriter);
		
		//BufferedWriter bw = new BufferedWriter(new FileWriter("lorem2.txt"));
		
		PrintStream ps = new PrintStream("lorem2.txt");
		
		ps.println("Hello, world!");
		ps.println("Bye, world.");
		ps.close();
		
		PrintWriter pw = new PrintWriter("lorem3.txt");
		pw.println("Hello, world!");
		pw.print("Bye, world.");
		pw.close();
	}

}
