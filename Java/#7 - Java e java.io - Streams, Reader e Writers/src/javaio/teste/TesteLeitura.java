package javaio.teste;

import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStreamReader;

public class TesteLeitura {

	public static void main(String[] args) throws IOException {

		FileInputStream fs = new FileInputStream("lorem.txt");
		InputStreamReader inReader = new InputStreamReader(fs);
		BufferedReader buffer = new BufferedReader(inReader);
		
		String linha = buffer.readLine();
		
		while(linha != null) {			
			System.out.println(linha);
			linha = buffer.readLine();
		}
		
		buffer.close();
	}

}
