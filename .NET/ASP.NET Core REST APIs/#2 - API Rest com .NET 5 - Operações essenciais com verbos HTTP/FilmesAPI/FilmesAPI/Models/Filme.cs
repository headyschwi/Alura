using System.ComponentModel.DataAnnotations;

namespace FilmesAPI.Models
{
    public class Filme
    {
        [Key][Required]
        public int id { get; set; }

        [Required(ErrorMessage = "O título é obrigatório.")]
        public string Titulo { get; set; }
        
        [Required(ErrorMessage = "O diretor é obrigatório.")]
        public string Diretor { get; set; }

        [Required(ErrorMessage = "O gênero é obrigatório.")]
        public string Genero { get; set; }
        
        [Required(ErrorMessage = "A duração é obrigatório.")] [Range(1,240, ErrorMessage = "O filme deve ter entre 1 e 240 minutos.")]
        public int Duracao { get; set; }

    }
}
