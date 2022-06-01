package javaio.teste;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.OutputStreamWriter;
import java.io.Writer;

public class TesteCopiaArquivo {
	public static void main(String[] args) throws IOException {
		InputStream in_fs = System.in; //new FileInputStream("lorem.txt")
		InputStreamReader in_reader = new InputStreamReader(in_fs);
		BufferedReader in_buffer = new BufferedReader(in_reader);

		OutputStream out_fs = System.out;  //new FileOutputStream("lorem2.txt");
		Writer out_writer = new OutputStreamWriter(out_fs);
		BufferedWriter out_buffer = new BufferedWriter(out_writer);

		String linha = in_buffer.readLine();

		while(linha != null && !linha.isEmpty()) {
			out_buffer.write(linha);
			out_buffer.newLine();
			out_buffer.flush();
			linha = in_buffer.readLine();		
		}

		in_buffer.close();
		out_buffer.close();
	}

}
